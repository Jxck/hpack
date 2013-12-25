package hpack

import (
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
