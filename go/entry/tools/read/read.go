package read

import (
	"os"
	p "path/filepath"

	"g.sacerb.com/sacer/go/entry/tools"
)

type FileInfo struct {
	Path     string
	DirEntry os.DirEntry
}

func (fi *FileInfo) IsDir() bool {
	return fi.DirEntry.IsDir()
}

func GetFiles(path string, withDirs bool) ([]*FileInfo, error) {
	l, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	list := []*FileInfo{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())

		if tools.IsDontIndex(filepath) {
			continue
		}

		if !withDirs && fi.IsDir() {
			continue
		}

		list = append(list, &FileInfo{
			Path:     filepath,
			DirEntry: fi,
		})
	}

	return list, nil
}
