package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	str := "hello"
	rs := Reverse(str)
	assert.Equal(t, "olleh", rs)
}
