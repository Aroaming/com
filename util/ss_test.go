package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSS(t *testing.T) {
	str := "hello"
	b := s2b(str)
	assert.Equal(t, []byte{'h', 'e', 'l', 'l', 'o'}, b)
	str = b2s(b)
	assert.Equal(t, "hello", str)
}
