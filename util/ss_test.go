package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSS(t *testing.T) {
	str := "hello"
	b := String2Byte(str)
	assert.Equal(t, []byte{'h', 'e', 'l', 'l', 'o'}, b)
	str = Byte2String(b)
	assert.Equal(t, "hello", str)

	type RoomRole struct {
		Room string
		User string
	}
	rr := RoomRole{
		Room: "123456",
		User: "abcdef",
	}

	data, err := StructToByte(rr)
	assert.Nil(t, err)
	var testRR RoomRole
	err = ByteToStruct(data, &testRR)
	assert.Nil(t, err)
	assert.Equal(t, "123456", testRR.Room)
	assert.Equal(t, "abcdef", testRR.User)
}
