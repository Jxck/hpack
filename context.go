package hpack

import (
	"bytes"
	"flag"
	"fmt"
	. "github.com/jxck/logger"
	"log"
	"net/http"
)

var verbose bool
var loglevel int

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose out")
	flag.IntVar(&loglevel, "l", 0, "log level (1 ERR, 2 WARNING, 3 INFO, 4 DEBUG)")
	flag.Parse()
	LogLevel(loglevel)
	Verbose(verbose)
}

type Context struct {
	HeaderTable  HeaderTable
	ReferenceSet ReferenceSet
	EmittedSet   EmittedSet
}

func NewRequestContext() *Context {
	var context = &Context{
		HeaderTable:  NewRequestHeaderTable(),
		ReferenceSet: NewReferenceSet(),
		EmittedSet:   NewEmittedSet(),
	}
	return context
}

func NewResponseContext() *Context {
	var context = &Context{
		HeaderTable:  NewResponseHeaderTable(),
		ReferenceSet: NewReferenceSet(),
		EmittedSet:   NewEmittedSet(),
	}
	return context
}

func (c *Context) Decode(wire []byte) {
	// EmittedSet を clean
	c.EmittedSet = NewEmittedSet()

	frames := Decode(wire)
	for _, frame := range frames {
		switch f := frame.(type) {
		case *IndexedHeader:
			// HT にあるエントリをそのまま使う
			header := c.HeaderTable.Headers[f.Index]
			Debug(fmt.Sprintf("%T: use %v(HT[%v])", f, header, f.Index))
			if header.Value == c.ReferenceSet[header.Name] {
				// refset にある場合は消す
				Debug(fmt.Sprintf("delete from refset (%q, %q)", header.Name, header.Value))
				c.ReferenceSet.Del(header.Name)
			} else {
				// refset にない場合は加える
				Debug(fmt.Sprintf("emit and add to refset (%q, %q)", header.Name, header.Value))
				c.EmittedSet.Emit(header.Name, header.Value)
				c.ReferenceSet.Add(header.Name, header.Value)
			}
		case *IndexedNameWithoutIndexing:
			// HT にある名前だけ使う
			// HT も refset も更新しない
			header := c.HeaderTable.Headers[f.Index]
			Debug(fmt.Sprintf("%T use name %v(HT[%v]) and value=%q", f, header.Name, f.Index, f.ValueString))
			Debug(fmt.Sprintf("emit (%q, %q)", header.Name, f.ValueString))
			c.EmittedSet.Emit(header.Name, f.ValueString)
		case *NewNameWithoutIndexing:
			// Name/Value ペアを送る
			// HT も refset も更新しない
			Debug(fmt.Sprintf("%T use name=%q and value=%q", f, f.NameString, f.ValueString))
			Debug(fmt.Sprintf("emit (%q, %q)", f.NameString, f.ValueString))
			c.EmittedSet.Emit(f.NameString, f.ValueString)
		case *IndexedNameWithIncrementalIndexing:
			// HT にある名前だけ使い
			// HT に新しく追記する
			// refset も更新する
			name := c.HeaderTable.Headers[f.Index].Name
			value := f.ValueString
			Debug(fmt.Sprintf("%T %v(HT[%v]) value=%q", f, name, f.Index, value))
			Debug(fmt.Sprintf("emit and add refeset, HT (%q, %q)", name, value))
			c.EmittedSet.Emit(name, value)
			c.HeaderTable.Add(name, value)
			c.ReferenceSet.Add(name, value)
		case *NewNameWithIncrementalIndexing:
			// Name/Value ペアを送る
			// HT と refset にも追記
			name, value := f.NameString, f.ValueString
			Debug(fmt.Sprintf("%T use name=%q and value=%q", f, name, value))
			Debug(fmt.Sprintf("emit and add refeset, HT (%q, %q)", name, value))
			c.EmittedSet.Emit(name, value)
			c.HeaderTable.Add(name, value)
			c.ReferenceSet.Add(name, value)
		case *IndexedNameWithSubstitutionIndexing:
			// HT[substituted index]  の 中身を
			// HT[index] と value で置き換える
			// refset も更新する
			name := c.HeaderTable.Headers[f.Index].Name
			value := f.ValueString
			Debug(fmt.Sprintf("%T name=%q value=%q", f, name, value))
			Debug(fmt.Sprintf("emit and add refeset, replace HT (%q, %q)", name, value))
			c.EmittedSet.Emit(name, value)
			c.HeaderTable.Replace(name, value, f.SubstitutedIndex)
			c.ReferenceSet.Add(name, value)
		case *NewNameWithSubstitutionIndexing:
			// HT[substituted index]  の 中身を
			// name と value で置き換える
			// refset も更新する
			name, value := f.NameString, f.ValueString
			Debug(fmt.Sprintf("%T HT[%v]=%v value=%q", f, f.Index, name, value))
			Debug(fmt.Sprintf("emit and add refeset, replace HT (%q, %q)", name, value))
			c.EmittedSet.Emit(name, value)
			c.HeaderTable.Replace(name, value, f.SubstitutedIndex)
			c.ReferenceSet.Add(name, value)
		default:
			log.Fatal("%T", f)
		}
	}
	// reference set の emitt されてないものを emit する
	for name, value := range c.ReferenceSet {
		if !c.EmittedSet.Check(name, value) {
			c.EmittedSet.Emit(name, value)
		}
	}
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
	for name, value := range c.ReferenceSet {
		if headerSet[name] != value {
			c.ReferenceSet.Del(name)

			// Header Table を探して、 index だけ取り出す
			index, _ := c.HeaderTable.SearchHeader(name, value)

			// Indexed Header を生成
			frame := CreateIndexedHeader(uint64(index))
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
			frame := CreateIndexedHeader(uint64(index))
			f := frame.Encode()
			Debug(fmt.Sprintf("indexed header index=%v", index))
			Debug(fmt.Sprintf("add to refset (%q, %q)", name, value))
			c.ReferenceSet.Add(name, value)
			buf.Write(f.Bytes())
		} else if index != -1 { // HT に name だけある
			// Indexed Name Without Indexing
			// value だけ送る。 HT は更新しない。
			frame := CreateIndexedNameWithoutIndexing(uint64(index), value)
			f := frame.Encode()
			Debug(fmt.Sprintf("literal header without indexing, name index=%v value=%q", index, value))
			buf.Write(f.Bytes())
		} else { // HT に name も value もない
			// New Name Without Indexing
			// name, value を送って HT は更新しない。
			frame := CreateNewNameWithoutIndexing(name, value)
			f := frame.Encode()
			Debug(fmt.Sprintf("literal header without indexing, new name name=%q value=%q", name, value))
			buf.Write(f.Bytes())
		}
	}
	return buf.Bytes()
}
