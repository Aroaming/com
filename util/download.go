package util

import (
	"io"
	"net/http"
	"os"

	"log"
)

//DownloadFile form url to name,just for a little file
func DownloadFile(url, name string) (err error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("download file url:%s error:%s", url, err)
		return
	}
	f, err := os.Create(name)
	if err != nil {
		log.Printf("create file error:%s", err)
		return
	}
	io.Copy(f, res.Body)
	return
}
