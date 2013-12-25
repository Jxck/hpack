package hpack

import (
	"flag"
	"fmt"
	. "github.com/jxck/color"
	. "github.com/jxck/logger"
	"log"
)

var verbose bool
var loglevel int

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose out")
	flag.IntVar(&loglevel, "l", 0, "log level (1 ERR, 2 WARNING, 3 INFO, 4 DEBUG)")
	flag.Parse()
	LogLevel(loglevel)
	Verbose(verbose)
	log.SetFlags(log.Lshortfile)
}

type Context struct {
	HT HeaderTable
	RS *ReferenceSet
}

func NewContext() Context {
	return Context{
		HT: HeaderTable{},
		RS: NewReferenceSet(),
	}
}

func (c *Context) Decode(wire []byte) {
	frames := Decode(wire)
	for _, frame := range frames {
		switch f := frame.(type) {
		case *IndexedHeader:
			index := int(f.Index)

			/**
			 * idx=0 の場合 Reference Set を空にする
			 */
			if index == 0 {
				c.RS.Empty()
			}

			/**
			 * idx が Reference Set にあった場合
			 * Reference Set から消す
			 */
			// if RefSet.Has(index) {
			//   RefSet.Remove(inx)
			//   continue
			// }

			/**
			 * idx が Reference Set に無い場合
			 * 該当のエントリを取り出す
			 */
			var headerField HeaderField

			if index > c.HT.Len() {
				// Static Header Table の中にあった場合
				index = index - c.HT.Len() - 1
				headerField = StaticHeaderTable[index]
				Debug(Red("== Indexed - Add =="))
				Debug(fmt.Sprintf("  idx = %v", index))
				Debug(fmt.Sprintf("-> ST[%v] = %v", index, headerField))

				// Emit
				Debug(Blue("Emit"))

				// ヘッダテーブルにコピーする
				// insertedIndex := c.HT.Push(headerField)

				// その参照を RefSet に追加する
				// RefSet.Add(insertedIndex)

			} else {
				// Header Table の中にあった場合
				Debug(Red("== Indexed - Add =="))
				Debug(fmt.Sprintf("  idx = %v", index))

				// Emit
				Debug(Blue("Emit"))

				// その参照を RefSet に追加する
				// RefSet.Add(index)
			}
		case *IndexedLiteral:
			log.Printf("%v", f)
			Debug(Red(fmt.Sprintf("== Literal Indexed (idx=%t) ==", f.Indexing)))

			if f.Indexing {
				// HT に追加する場合

				// Emit
				Debug(Blue("Emit"))

				// ヘッダテーブルにコピーする
				// insertedIndex := c.HT.Push(headerField)

				// その参照を RefSet に追加する
				// RefSet.Add(insertedIndex)

			} else {
				// HT に追加しない場合

				// Emit
				Debug(Blue("Emit"))
			}
		case *StringLiteral:
			log.Printf("%v", f)
			Debug(Red(fmt.Sprintf("== Literal Indexed (idx=%t) ==", f.Indexing)))

			if f.Indexing {
				// HT に追加する場合

				// Emit
				Debug(Blue("Emit"))

				// ヘッダテーブルにコピーする
				// insertedIndex := c.HT.Push(headerField)

				// その参照を RefSet に追加する
				// RefSet.Add(insertedIndex)

			} else {
				// HT に追加しない場合

				// Emit
				Debug(Blue("Emit"))
			}

		default:
			log.Fatal("%T", f)
		}
	}
	// reference set の emitt されてないものを emit する
	//for name, value := range c.ReferenceSet {
	//	if !c.EmittedSet.Check(name, value) {
	//		c.EmittedSet.Emit(name, value)
	//	}
	//}
}

/*
func (c *Context) Encode(header http.Header) []byte {
	var buf bytes.Buffer

	// http.Header を HeaderSet に変換
	headerSet := HeaderToHeaderSet(header)

	// ReferenceSet の中から消すべき値を消す
	buf.Write(c.CleanReferenceSet(headerSet))

	// Header Set の中から送らない値を消す
	c.CleanHeaderSet(headerSet)

	// Header Table にあるやつを処理
	buf.Write(c.ProcessHeader(headerSet))
	return buf.Bytes()
}

// 1. 不要なエントリを reference set から消す
func (c *Context) CleanReferenceSet(headerSet HeaderSet) []byte {
	var buf bytes.Buffer
	// reference set の中にあって、 header set の中に無いものは
	// 相手の reference set から消さないといけないので、
	// indexed representation でエンコードして
	// reference set からは消す
	for name, value := range c.ReferenceSet {
		if headerSet[name] != value {
			c.ReferenceSet.Del(name)

			// Header Table を探して、 index だけ取り出す
			index, _ := c.HeaderTable.SearchHeader(name, value)

			// Indexed Header を生成
			frame := NewIndexedHeader(uint64(index))
			f := frame.Encode()
			buf.Write(f.Bytes())
			Debug(fmt.Sprintf("indexed header index=%v removal from reference set", index))
		}
	}
	return buf.Bytes()
}

// 2. 送る必要の無いものを header set から消す
func (c *Context) CleanHeaderSet(headerSet HeaderSet) {
	for name, value := range c.ReferenceSet {
		if headerSet[name] == value {
			delete(headerSet, name)
			// TODO: "common-header" としてマーク
			Debug(fmt.Sprintf("remove from header set %v %v", name, value))
		}
	}
}

// 3 と 4. 残りの処理
func (c *Context) ProcessHeader(headerSet HeaderSet) []byte {
	var buf bytes.Buffer
	for name, value := range headerSet {
		index, h := c.HeaderTable.SearchHeader(name, value)
		if h != nil { // 3.1 HT にエントリがある
			// Indexed Heaer で index だけ送れば良い
			frame := NewIndexedHeader(uint64(index))
			f := frame.Encode()
			Debug(fmt.Sprintf("indexed header index=%v", index))
			Debug(fmt.Sprintf("add to refset (%q, %q)", name, value))
			c.ReferenceSet.Add(name, value)
			buf.Write(f.Bytes())
		} else if index != -1 { // HT に name だけある
			// Indexed Name Without Indexing
			// value だけ送る。 HT は更新しない。
			frame := NewIndexedNameWithoutIndexing(uint64(index), value)
			f := frame.Encode()
			Debug(fmt.Sprintf("literal header without indexing, name index=%v value=%q", index, value))
			buf.Write(f.Bytes())
		} else { // HT に name も value もない
			// New Name Without Indexing
			// name, value を送って HT は更新しない。
			frame := NewNewNameWithoutIndexing(name, value)
			f := frame.Encode()
			Debug(fmt.Sprintf("literal header without indexing, new name name=%q value=%q", name, value))
			buf.Write(f.Bytes())
		}
	}
	return buf.Bytes()
}
*/
