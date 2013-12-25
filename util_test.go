package hpack

import (
	"reflect"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	fixture := []string{":foo", ":foo:", "::foo", ":foo:foo", "bar"}
	expected := []string{"foo", "foo:", "foo", "foo:foo", "bar"}
	actual := make([]string, len(fixture))
	for i, name := range fixture {
		actual[i] = RemovePrefix(name)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
