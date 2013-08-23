package hpac

import (
	"bytes"
	"log"
	"testing"
)

/*
0x44      (literal header with incremental indexing, name index = 3)
0x16      (header value string length = 22)
/my-example/index.html
0x4D      (literal header with incremental indexing, name index = 12)
0x0D      (header value string length = 13)
my-user-agent
0x40      (literal header with incremental indexing, new name)
0x0B      (header name string length = 11)
mynewheader
0x05      (header value string length = 5)
first
*/

func TestDecoder(t *testing.T) {
	log.SetFlags(log.Lshortfile)

	buf := bytes.NewBuffer([]byte{0x44, 0x16}) // 0100 0100 0001 0110
	log.Println(buf.WriteString("/my-example/index.html"))
	frame := DecodeHeader(buf)

	if frame.Flag1 != 0 {
		t.Errorf("got %v\nwant %v", frame.Flag1, 0)
	}
	if frame.Flag2 != 1 {
		t.Errorf("got %v\nwant %v", frame.Flag2, 1)
	}
	if frame.Flag3 != 0 {
		t.Errorf("got %v\nwant %v", frame.Flag3, 0)
	}
	// TODO: Pending
	// if frame.Index != 3 {
	// 	t.Errorf("got %v\nwant %v", frame.Index, 3)
	// }
	if frame.ValueLength != 22 {
		t.Errorf("got %v\nwant %v", frame.Index, 22)
	}
	if frame.ValueString != "/my-example/index.html" {
		t.Errorf("got %v\nwant %v", frame.ValueString, "/my-example/index.html")
	}
}
