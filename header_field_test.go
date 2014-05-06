package hpack

import (
	assert "github.com/jxck/assertion"
	"testing"
)

func TestHeaderFieldSize(t *testing.T) {
	h := NewHeaderField("hello", "world")
	assert.Equal(t, h.Size(), 42)
}
