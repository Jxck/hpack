package hpac

import (
	"testing"
)

func TestReferenceSet(t *testing.T) {
	ref := ReferenceSet{}

	ref.Add("hoge", "fuga")
	if len(ref) != 1 {
		t.Errorf("got %v\nwant %v", len(ref), 1)
	}
	if ref["hoge"] != "fuga" {
		t.Errorf("got %v\nwant %v", ref["hoge"], "fuga")
	}

	ref.Set("hoge", "piyo")
	if len(ref) != 1 {
		t.Errorf("got %v\nwant %v", len(ref), 1)
	}
	if ref["hoge"] != "piyo" {
		t.Errorf("got %v\nwant %v", ref["hoge"], "piyo")
	}

	ref.Del("hoge")
	if len(ref) != 0 {
		t.Errorf("got %v\nwant %v", len(ref), 0)
	}
}
