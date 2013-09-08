package hpac

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ref := ReferenceSet{}

	ref.Add("hoge", "fuga")
	if len(ref) != 1 {
		t.Errorf("got %v\nwant %v", len(ref), 1)
	}
	if ref["hoge"] != "fuga" {
		t.Errorf("got %v\nwant %v", ref["hoge"], "fuga")
	}
}

func TestDel(t *testing.T) {
	ref := ReferenceSet{"hoge": "fuga"}

	ref.Del("hoge")
	if len(ref) != 0 {
		t.Errorf("got %v\nwant %v", len(ref), 0)
	}
}
