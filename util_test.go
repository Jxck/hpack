package hpac

import (
	"testing"
)

func TestCompareSlice(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"a", "b"}

	if CompareSlice(a, b) != true {
		t.Fatal("slices are same", a, b)
	}

	a = []string{"a", "b"}
	b = []string{"a"}

	if CompareSlice(a, b) != false {
		t.Fatal("slices are not same", a, b)
	}
}
