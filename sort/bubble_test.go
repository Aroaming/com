package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	nums := []int{4, 3, 2, 1, 5}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, bubble(nums))
}
