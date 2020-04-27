package hash

import (
	"crypto/md5"
	"crypto/sha256"
)

func hash() string {
	var data []byte
	md5.Sum(data)
	sha256.Sum224()
	return ""
}
