package disk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskinfo(t *testing.T) {
	info := "/dev/nvme0n1p2 ext4 244568380 149105828 82969496   65% /\n"
	afterInfo := trimLine(info)
	assert.Equal(t, "/dev/nvme0n1p2 ext4 244568380 149105828 82969496 65% /", afterInfo)
}
