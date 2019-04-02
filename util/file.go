package util

import (
	"os"
)

func IsFile(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

//whether a file or dir
func IsExsit(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
