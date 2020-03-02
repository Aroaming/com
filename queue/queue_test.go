package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := New()
	queue.Put("1")
	queue.Put("2")
	queue.Put("3")
	assert.Equal(t, 3, queue.Size())

	queue.Cut()
	assert.Equal(t, 2, queue.Size())

	queue.Put("4")
	queue.Put("5")
	queue.Put("6")
	value := queue.Get(2)
	assert.Equal(t, "3", value)

	queue.Cut()
	assert.Equal(t, 4, queue.Size())
	queue.Clear()
	assert.Equal(t, 0, queue.Size())
	queue.Put("1")
	assert.Equal(t, 1, queue.Size())
}
