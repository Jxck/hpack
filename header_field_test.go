package hpack

import (
	assert "github.com/Jxck/assertion"
	"testing"
)

func TestHeaderFieldSize(t *testing.T) {
	h := NewHeaderField("hello", "world")
	assert.Equal(t, h.Size(), 42)
}
