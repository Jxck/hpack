package hpack

import (
	"flag"
	"fmt"
	. "github.com/jxck/color"
	. "github.com/jxck/logger"
	"github.com/jxck/swrap"
	"log"
)

var loglevel int

func init() {
	flag.IntVar(&loglevel, "l", 0, "log level (1 ERR, 2 WARNING, 3 NOTICE, 4 INFO, 5 DEBUG)")
	flag.Parse()
	LogLevel(loglevel)
	log.SetFlags(log.Lshortfile)
}

type Context struct {
	HT *HeaderTable
	RS *ReferenceSet
	ES *HeaderSet
}

func NewContext(SETTINGS_HEADER_TABLE_SIZE int) *Context {
	return &Context{
		HT: NewHeaderTable(SETTINGS_HEADER_TABLE_SIZE),
		RS: NewReferenceSet(),
		ES: NewHeaderSet(),
	}
}

func (c *Context) Decode(wire []byte) {
	// 各デコードごとに前回のをリセットする。
	c.ES = NewHeaderSet()
	c.RS.Reset()
	Debug(Red("clean Emitted Set"))
	Trace(Cyan(
		"\n===== Before Decode =====")+
		"%v"+Cyan(
		"==========================="),
		c.Dump())

	frames := Decode(wire)
	for _, frame := range frames {
		switch f := frame.(type) {
		case *IndexedHeader:
			index := int(f.Index)

			if index == 0 {
				/**
				 * idx=0 の場合 Option を見る
				 */

				if f.Option == 128 {
					/**
					 * Reference Set Emptying
					 */
					Debug(Red("Reference Set Emptying"))
					c.RS.Empty()
				} else if f.Option < 128 {
					/**
					 * Maximum Header Table Size Change
					 */
					Debug(Red("Maximum Header Table Size Change"))
					// TODO: change header table size
				}

				continue
			}

			var headerField *HeaderField

			if index > c.HT.Len() {
				/**
				 * Static Header Table の中にある場合
				 */
				i := index - c.HT.Len() - 1
				headerField = StaticHeaderTable[i]

				if c.RS.Has(headerField) {
					/**
					 * 参照が Reference Set にあった場合
					 * Reference Set から消す
					 */

					Debug(Red(fmt.Sprintf("== Indexed - Remove ==")))
					Debug(fmt.Sprintf("  idx = %v", index))
					Debug(fmt.Sprintf("  -> HT[%v] = %v", index, headerField))

					// remove
					c.RS.Remove(headerField)
				} else {
					/**
					* 参照が Reference Set に無い場合
					* 該当のエントリを取り出す
					 */

					Debug(Red("== Indexed - Add =="))
					Debug(fmt.Sprintf("  idx = %v", index))
					Debug(fmt.Sprintf("  -> ST[%v] = %v", index, headerField))

					// Emit
					Debug(Blue("\tEmit"))
					c.ES.Emit(headerField)

					// ヘッダテーブルにコピーする
					Debug(Blue("\tAdd to HT"))
					c.Push(headerField)

					// その参照を RefSet に追加する
					Debug(Blue("\tAdd to RS"))
					c.RS.Add(headerField, EMITTED)
				}
			} else {
				/**
				 * Header Table の中にある場合
				 */

				// 実態は配列なので 0 オリジン
				i := index - 1
				headerField = c.HT.HeaderFields[i]

				if c.RS.Has(headerField) {
					/**
					 * 参照が Reference Set にあった場合
					 * Reference Set から消す
					 */
					Debug(Red(fmt.Sprintf("== Indexed - Remove ==")))
					Debug(fmt.Sprintf("  idx = %v", index))
					Debug(fmt.Sprintf("  -> HT[%v] = %v", index, headerField))

					// remove
					c.RS.Remove(headerField)
				} else {
					/**
					* 参照が Reference Set に無い場合
					 */
					Debug(Red("== Indexed - Add =="))
					Debug(fmt.Sprintf("  idx = %v", index))
					Debug(fmt.Sprintf("  -> HT[%v] = %v", index, headerField))

					// Emit
					Debug(Blue("\tEmit"))
					c.ES.Emit(headerField)

					// その参照を RefSet に追加する
					Debug(Blue("\tAdd to RS"))
					c.RS.Add(headerField, EMITTED)
				}
			}
		case *IndexedLiteral:

			// Index 先の Name と Literal Value から HeaderField を生成
			index := int(f.Index)
			var name, value string

			if index > c.HT.Len() {
				/**
				 * Static Header Table の中にある場合
				 */
				i := index - c.HT.Len() - 1
				name = StaticHeaderTable[i].Name
			} else {
				/**
				 * Header Table の中にある場合
				 */
				i := index - 1
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

			if f.Indexing {
				/**
				 * HT に追加する場合
				 */

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)

				// ヘッダテーブルにコピーする
				Debug(Blue("\tAdd to HT"))
				c.Push(headerField)

				// その参照を RefSet に追加する
				Debug(Blue("\tAdd to RS"))
				c.RS.Add(headerField, EMITTED)

			} else {
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
			if f.Indexing {
				// HT に追加する場合

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)

				// ヘッダテーブルにコピーする
				Debug(Blue("\tAdd to HT"))
				c.Push(headerField)

				// その参照を RefSet に追加する
				Debug(Blue("\tAdd to RS"))
				c.RS.Add(headerField, EMITTED)

			} else {
				// HT に追加しない場合

				// Emit
				Debug(Blue("\tEmit"))
				c.ES.Emit(headerField)
			}

		default:
			log.Fatal("%T", f)
		}
	}

	// reference set の残りを全て emit する
	for _, referencedField := range *c.RS {
		if !referencedField.Emitted {
			headerField := referencedField.HeaderField
			Debug(Blue("\tEmit rest entries ")+"%v", *headerField)
			c.ES.Emit(headerField)
		}
	}
}

// removing entry from top
// until make space of size in Header Table
// Evict された参照を RS からも消すために、 Context の方でやる。
func (c *Context) Eviction() {
	for c.HT.Size() > c.HT.HEADER_TABLE_SIZE {
		// サイズが収まるまで減らす
		Debug(Red("Eviction")+" %v", c.HT.HeaderFields[len(c.HT.HeaderFields)-1])
		removed := c.HT.Remove(len(c.HT.HeaderFields) - 1)

		// 消したエントリへの参照を RS からも消す
		c.RS.Remove(removed)
	}
	return
}

// Push new enctory to Header Table
// and Eviction
func (c *Context) Push(hf *HeaderField) {
	c.HT.Push(hf)
	c.Eviction()
}

// Dump for Debug
func (c *Context) Dump() string {
	return fmt.Sprintf("%v%v%v", c.HT.Dump(), c.RS.Dump(), c.ES.Dump())
}

func (c *Context) Encode(headerSet HeaderSet) []byte {
	var buf swrap.SWrap

	// ReferenceSet を空にする
	indexHeader := NewIndexedHeader(0)
	indexHeader.Option = 0x80

	emptyRef := *indexHeader.Encode()
	buf.Merge(emptyRef)

	// 全て StringLiteral(Indexing = false) でエンコード
	for _, h := range headerSet {
		sl := NewStringLiteral(false, h.Name, h.Value)
		buf.Merge(*sl.EncodeHuffman())
	}

	return buf.Bytes()
}
