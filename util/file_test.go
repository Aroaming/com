package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	assert.Equal(t, "100B", HumanFileSize(100))
}
