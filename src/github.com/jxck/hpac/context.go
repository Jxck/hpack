package hpac

import (
	"bytes"
	"log"
	"net/http"
)

type Context struct {
	requestHeaderTable  HeaderTable
	responseHeaderTable HeaderTable
	referenceSet        ReferenceSet
}

func NewContext() *Context {
	var context = &Context{
		requestHeaderTable:  NewRequestHeaderTable(),
		responseHeaderTable: NewResponseHeaderTable(),
		referenceSet:        ReferenceSet{},
	}
	return context
}

func (c *Context) Decode(wire []byte) {
	frames := Decode(wire)
	for _, frame := range frames {
		switch f := frame.(type) {
		case *IndexedHeader:
			log.Printf("%T index=%v", f, f.Index)
			header := c.requestHeaderTable[f.Index]
			log.Printf("HT[%v] = %v", f.Index, header)

			if header.Value == c.referenceSet[header.Name] {
				log.Println("exist in refset")
				c.referenceSet.Del(header.Name)
			} else {
				log.Println("no exist in refet")
				log.Println("TODO: emit") // TODO:emit
				c.referenceSet[header.Name] = header.Value
			}
		case *IndexedNameWithIncrementalIndexing:
			log.Printf("%T index=%v value=%q", f, f.Index, f.ValueString)
			header := c.requestHeaderTable[f.Index]
			log.Printf("HT[%v] = %v", f.Index, header)
			log.Println("TODO: emit") // TODO:emit
			c.requestHeaderTable[f.Index].Value = f.ValueString
			c.referenceSet[header.Name] = f.ValueString
		case *NewNameWithoutIndexing:
			log.Printf("%T name=%q value=%q", f, f.NameString, f.ValueString)
			c.referenceSet[f.NameString] = f.ValueString
		default:
			log.Printf("%T", f)
		}
	}
	log.Printf("refset: %v", c.referenceSet)
	log.Printf("HT: %v", c.requestHeaderTable)
}

func (c *Context) Encode(header http.Header) []byte {
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
	for name, value := range c.referenceSet {
		if headerSet[name] != value {
			log.Println("remove from refset", name, value)
			c.referenceSet.Del(name)

			// Header Table を探して、 index だけ取り出す
			index, _ := c.requestHeaderTable.SearchHeader(name, value)

			// Indexed Header を生成
			frame := CreateIndexedHeader(uint64(index))
			f := EncodeHeader(frame)
			buf.Write(f.Bytes())
		}
	}
	return buf.Bytes()
}

// 2. 送る必要の無いものを header set から消す
func (c *Context) CleanHeaderSet(headerSet HeaderSet) {
	for name, value := range c.referenceSet {
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
		index, h := c.requestHeaderTable.SearchHeader(name, value)
		if h != nil { // 3.1 HT にエントリがある
			// Indexed Heaer で index だけ送れば良い
			frame := CreateIndexedHeader(uint64(index))
			f := EncodeHeader(frame)
			log.Printf("indexed header {%v:%v} is in HT[%v] (%v)", name, value, index, f.Bytes())
			buf.Write(f.Bytes())
		} else if index != -1 { // HT に name だけある
			// Indexed Name With Incremental Indexing
			// value だけ送って、 HT にエントリを追加する。
			frame := CreateIndexedNameWithIncrementalIndexing(uint64(index), value)
			f := EncodeHeader(frame)
			log.Printf("literal with index {%v:%v} is in HT[%v] (%v)", name, value, index, f.Bytes())
			buf.Write(f.Bytes())
		} else { // HT に name も value もない
			// New Name Without Indexing
			// name, value を送って HT は変えない
			frame := CreateNewNameWithoutIndexing(name, value)
			f := EncodeHeader(frame)
			log.Printf("literal without index {%v:%v} is not in HT (%v)", name, value, f.Bytes())
			buf.Write(f.Bytes())
		}
	}
	return buf.Bytes()
}
