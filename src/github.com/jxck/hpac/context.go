package hpac

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type Context struct {
	RequestHeaderTable  HeaderTable
	ResponseHeaderTable HeaderTable
	ReferenceSet        ReferenceSet
	EmittedSet          http.Header
}

func NewContext() *Context {
	var context = &Context{
		RequestHeaderTable:  NewRequestHeaderTable(),
		ResponseHeaderTable: NewResponseHeaderTable(),
		ReferenceSet:        ReferenceSet{},
		EmittedSet:          http.Header{},
	}
	return context
}

func (c *Context) Decode(wire []byte) {
	fmt.Println("Decode")
	// EmittedSet を clean
	c.EmittedSet = http.Header{}

	frames := Decode(wire)
	for _, frame := range frames {
		switch f := frame.(type) {
		case *IndexedHeader:
			// HT にあるエントリをそのまま使う
			header := c.RequestHeaderTable[f.Index]
			log.Printf("%T HT[%v] = %v", f, f.Index, header)

			if header.Value == c.ReferenceSet[header.Name] {
				// refset にある場合は消す
				log.Printf("delete from refset (%q, %q)", header.Name, header.Value)
				c.ReferenceSet.Del(header.Name)
			} else {
				// refset にない場合は加える
				log.Printf("emit and add to refset (%q, %q)", header.Name, header.Value)
				c.EmittedSet.Add(header.Name, header.Value)
				c.ReferenceSet.Add(header.Name, header.Value)
			}
		case *IndexedNameWithoutIndexing:
			// HT にある名前だけ使う
			header := c.RequestHeaderTable[f.Index]
			log.Printf("%T HT[%v] = %v value=%q", f, f.Index, header.Name, f.ValueString)

			// without indexing なので refset には入れない
			log.Printf("emit (%q, %q)", header.Name, f.ValueString)
			c.EmittedSet.Add(header.Name, f.ValueString)
		case *NewNameWithoutIndexing:
			// Name/Value ペアを送る
			// HT も refset も更新しない
			log.Printf("%T name=%q value=%q", f, f.NameString, f.ValueString)
			log.Printf("emit (%q, %q)", f.NameString, f.ValueString)
			c.EmittedSet.Add(f.NameString, f.ValueString)
		case *IndexedNameWithIncrementalIndexing:
			// HT にある名前だけ使い
			// HT に新しく追記する
			// refset も更新する
			header := c.RequestHeaderTable[f.Index]
			name, value := header.Name, f.ValueString
			log.Printf("emit and add refeset, HT (%q, %q)", name, value)
			c.EmittedSet.Add(name, value)
			c.RequestHeaderTable.Add(name, value)
			c.ReferenceSet.Add(name, value)
		case *NewNameWithIncrementalIndexing:
		case *IndexedNameWithSubstitutionIndexing:
		case *NewNameWithSubstitutionIndexing:
		default:
			log.Fatal("%T", f)
		}
	}
	// reference set の emitt されてないものを emit する
	for name, value := range c.ReferenceSet {
		if c.EmittedSet.Get(name) != value {
			c.EmittedSet.Add(name, value)
		}
	}
}

func (c *Context) Encode(header http.Header) []byte {
	fmt.Println("Encode")
	var buf bytes.Buffer

	// http.Header を HeaderSet に変換
	headerSet := NewHeaderSet(header)

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
			index, _ := c.RequestHeaderTable.SearchHeader(name, value)

			// Indexed Header を生成
			frame := CreateIndexedHeader(uint64(index))
			f := frame.Encode()
			buf.Write(f.Bytes())

			log.Printf("indexed header index=%v removal from reference set", index)

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
			log.Println("remove from header set", name, value)
		}
	}
}

// 3 と 4. 残りの処理
func (c *Context) ProcessHeader(headerSet HeaderSet) []byte {
	var buf bytes.Buffer
	for name, value := range headerSet {
		index, h := c.RequestHeaderTable.SearchHeader(name, value)
		if h != nil { // 3.1 HT にエントリがある
			// Indexed Heaer で index だけ送れば良い
			frame := CreateIndexedHeader(uint64(index))
			f := frame.Encode()
			log.Printf("indexed header index=%v", index)
			log.Printf("add to refset (%q, %q)", name, value)
			c.ReferenceSet.Add(name, value)
			buf.Write(f.Bytes())
		} else if index != -1 { // HT に name だけある
			// Indexed Name Without Indexing
			// value だけ送る。 HT は更新しない。
			frame := CreateIndexedNameWithoutIndexing(uint64(index), value)
			f := frame.Encode()
			log.Printf("literal header without indexing, name index=%v value=%q", index, value)
			buf.Write(f.Bytes())
		} else { // HT に name も value もない
			// New Name Without Indexing
			// name, value を送って HT は更新しない。
			frame := CreateNewNameWithoutIndexing(name, value)
			f := frame.Encode()
			log.Printf("literal header without indexing, new name name=%q value=%q", name, value)
			buf.Write(f.Bytes())
		}
	}
	return buf.Bytes()
}
