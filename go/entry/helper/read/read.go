package read

import (
	"io/ioutil"
	"os"
	p "path/filepath"
	"stferal/go/entry/helper"
)

type FileInfo struct {
	Path     string
	FileInfo os.FileInfo
}

func (fi *FileInfo) IsDir() bool {
	return fi.FileInfo.IsDir()
}

func GetFiles(path string, withDirs bool) ([]*FileInfo, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	list := []*FileInfo{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())

		if helper.IsDontIndex(filepath) {
			continue
		}

		if !withDirs && fi.IsDir() {
			continue
		}

		list = append(list, &FileInfo{
			Path:     filepath,
			FileInfo: fi,
		})
	}

	return list, nil
}
