package hpack

import (
	"reflect"
	"testing"
)

func TestRemovePrefix(t *testing.T) {
	var params = []struct {
		fixture, expected string
	}{
		{":foo", "foo"},
		{":foo:", "foo:"},
		{"::foo", "foo"},
		{":foo:foo", "foo:foo"},
		{"bar", "bar"},
	}

	for _, param := range params {
		if RemovePrefix(param.fixture) != param.expected {
			t.Errorf("got %v\nwant %v", RemovePrefix(param.fixture), param.expected)
		}
	}
}

func TestHeaderToHeaderSet(t *testing.T) {
	actual := HeaderToHeaderSet(header)
	expected := headerSet
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
