package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNet(t *testing.T) {
	mac1 := "00:e0:4c:8f:c9:43"
	assert.Equal(t, true, CheckMacaddr(mac1))

	mac2 := "00:e0:4c:8f:c9:4p"
	assert.Equal(t, false, CheckMacaddr(mac2))
}
