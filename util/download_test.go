package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	err := DownloadFile("https://github.com/mymmsc/books/raw/master/%E9%A1%B9%E7%9B%AE%E7%AE%A1%E7%90%86/%E6%95%8F%E6%8D%B7%E5%BC%80%E5%8F%91.pdf", "敏捷开发.pdf")

	assert.NoError(t, err)
}
