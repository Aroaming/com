package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
)

//if file exsit return true
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

//file path not exsit,will be create dir.
//write data to file
func WriteFile(fileName string, data []byte) error {
	os.MkdirAll(path.Dir(fileName), os.ModePerm)
	return ioutil.WriteFile(fileName, data, 0655)
}

//Copy file from soure to target path
func Copy(src, dest string) error {
	//Gather file infomation to set back later
	sf, err := os.Lstat(src)
	if err != nil {
		return err
	}

	//symbolic link
	//
	if sf.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		//NOTE:os.Chmod and os.Chtimes don't recoganize symbolic link
		//which will lead "no such file or directory"
		return os.Symlink(target, dest)
	}
	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	if _, err = io.Copy(dw, sr); err != nil {
		return err
	}
	//set back file information
	if err = os.Chtimes(dest, sf.ModTime(), sf.ModTime()); err != nil {
		return err
	}
	return os.Chmod(dest, sf.Mode())
}

//return file size in bytes and possible error
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

//return file modified time and possible error
func FileMTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

func HumanFileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}
	return fmt.Sprintf(f+"%s", val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}
