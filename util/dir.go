package util

import (
	"errors"
	"os"
	"path"
	"strings"
)

//return ture if path is directory
//or return false if directort not exsit or it'a file
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func StatDir(rootPath string, includeDir ...bool) ([]string, error) {
	if !IsDir(rootPath) {
		return nil, errors.New("not a directory:" + rootPath)
	}

	isIncludeDir := false
	if len(includeDir) >= 1 {
		isIncludeDir = includeDir[0]
	}
	return statDir(rootPath, "", isIncludeDir, false, false)
}

func statDir(dirPath, recPath string, includeDir, isDIrOnly, followSymlinks bool) ([]string, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fis, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	statList := make([]string, 0)
	for _, fi := range fis {
		if strings.Contains(fi.Name(), ".DS_Store") {
			continue
		}
		relPath := path.Join(recPath, fi.Name())
		curPath := path.Join(dirPath, fi.Name())
		if fi.IsDir() {
			if includeDir {
				statList = append(statList, recPath+"/")
			}
			s, err := statDir(curPath, relPath, includeDir, isDIrOnly, followSymlinks)
			if err != nil {
				return nil, err
			}
			statList = append(statList, s...)
		} else if !isDIrOnly {
			statList = append(statList, s...)
		} else if followSymlinks && fi.Mode()&os.ModeSymlink != 0 {
			link, err := os.Readlink(curPath)
			if err != nil {
				return nil, err
			}
			if IsDir(link) {
				if includeDir {
					statList = append(statList, "/")
				}
				s, err := statDir(curPath, relPath, includeDir, isDIrOnly, followSymlinks)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return statList, nil
}
