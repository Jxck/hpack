package hpack

import (
	"flag"
	"fmt"
	. "github.com/Jxck/color"
	. "github.com/Jxck/logger"
	"github.com/Jxck/swrap"
	"log"
)

var loglevel int

var STATIC_HEADER_TABLE_SIZE = len(StaticHeaderTable)

func init() {
	flag.IntVar(&loglevel, "l", 0, "log level (1 ERR, 2 WARNING, 3 NOTICE, 4 INFO, 5 DEBUG, 6 TRACE)")
	flag.Parse()
	LogLevel(loglevel)
	log.SetFlags(log.Lshortfile)
}

type Context struct {
	HT *HeaderTable
	ES *HeaderList
}

func NewContext(SETTINGS_HEADER_TABLE_SIZE uint64) *Context {
	return &Context{
		HT: NewHeaderTable(SETTINGS_HEADER_TABLE_SIZE),
		ES: NewHeaderList(),
	}
}

func (c *Context) Decode(wire []byte) {
	// 各デコードごとに前回のをリセットする。
	c.ES = NewHeaderList()
	Debug(Red("clean Emitted Set"))
	Trace(Cyan(
		"\n===== Before Decode =====")+
		"%v"+Cyan(
		"==========================="),
		c.String())

	frames := Decode(wire)
	for _, frame := range frames {
		switch f := frame.(type) {
		case *IndexedHeader:
			index := int(f.Index)

			if index == 0 {
				// TODO: Decoding Error
				// log.Fatal("Decoding Error: The index value of 0 is not used.")
			}

			var headerField *HeaderField

			if index < STATIC_HEADER_TABLE_SIZE {
				/**
				 * Static Header Table の中にある場合
				 */
				// 実態は配列なので 0 オリジン
				i := index - 1
				headerField = &StaticHeaderTable[i]

				// Emit
				Debug(Red(fmt.Sprintf("== Indexed - Add ==")))
				Debug(fmt.Sprintf("  idx = %v", index))
				Debug(fmt.Sprintf("  -> ST[%v] = %v", i, headerField))
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)
			} else {
				/**
				 * Header Table の中にある場合
				 */

				// 実態は配列なので 0 オリジン
				i := index - STATIC_HEADER_TABLE_SIZE - 1
				headerField = c.HT.HeaderFields[i]

				/**
				* 参照が Reference Set に無い場合
				 */
				Debug(Red("== Indexed - Add =="))
				Debug(fmt.Sprintf("  idx = %v", index))
				Debug(fmt.Sprintf("  -> HT[%v] = %v", index, headerField))

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)
			}
		case *IndexedLiteral:

			// Index 先の Name と Literal Value から HeaderField を生成
			index := int(f.Index)
			var name, value string

			if index < STATIC_HEADER_TABLE_SIZE {
				/**
				 * Static Header Table の中にある場合
				 */
				i := index - 1
				name = StaticHeaderTable[i].Name
			} else {
				/**
				 * Header Table の中にある場合
				 */
				i := index - STATIC_HEADER_TABLE_SIZE - 1
				name = c.HT.HeaderFields[i].Name
			}

			value = f.ValueString

			// Header Field 生成
			headerField := NewHeaderField(name, value)

			Debug(Red("== Indexed Literal =="))
			Debug(fmt.Sprintf("  Indexed name (idx = %v)", index))
			Debug(fmt.Sprintf("  -> ST[%v].Name = %v", index, name))
			Debug(fmt.Sprintf("  Literal value (len = %v)", f.ValueLength))
			Debug(fmt.Sprintf("  %v", f.ValueString))

			switch f.Indexing {
			case WITH:
				/**
				 * HT に追加する場合
				 */

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)

				// ヘッダテーブルにコピーする
				Debug(Blue("\tAdd to HT"))
				c.Push(headerField)

			case WITHOUT:
				/**
				 * HT に追加しない場合
				 */

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)
			}

		case *StringLiteral:
			Debug(Red(fmt.Sprintf("== String Literal (%v) ==", f)))

			headerField := NewHeaderField(f.NameString, f.ValueString)
			switch f.Indexing {
			case WITH:
				// HT に追加する場合

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)

				// ヘッダテーブルにコピーする
				Debug(Blue("\tAdd to HT"))
				c.Push(headerField)

			case WITHOUT:
				// HT に追加しない場合

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)
			}
		case *ChangeHeaderTableSize:
			/**
			 * Maximum Header Table Size Change
			 */
			Debug(Red("Maximum Header Table Size Change"))
			// TODO: change header table size
		default:
			log.Fatal("%T", f)
		}
	}
}

func (c *Context) ChangeSize(size uint64) {
	c.HT.HEADER_TABLE_SIZE = size
	c.Eviction()
}

// removing entry from top
// until make space of size in Header Table
func (c *Context) Eviction() {
	for c.HT.Size() > c.HT.HEADER_TABLE_SIZE {
		// サイズが収まるまで減らす
		Debug(Red("Eviction")+" %v", c.HT.HeaderFields[len(c.HT.HeaderFields)-1])
		removed := c.HT.Remove(len(c.HT.HeaderFields) - 1)
		Debug(Yellow(fmt.Sprintf("Removed while Eviction: %v", removed)))
	}
	return
}

// Push new enctory to Header Table
// and Eviction
func (c *Context) Push(hf *HeaderField) {
	c.HT.Push(hf)
	c.Eviction()
}

// String for Debug
func (c *Context) String() string {
	return fmt.Sprintf("%v%v", c.HT.String(), c.ES.String())
}

func (c *Context) Encode(headerList HeaderList) []byte {
	var buf swrap.SWrap

	// 全て StringLiteral(Indexing = false) でエンコード
	for _, h := range headerList {
		sl := NewStringLiteral(WITHOUT, h.Name, h.Value)
		buf.Merge(*sl.EncodeHuffman())
	}

	return buf.Bytes()
}
