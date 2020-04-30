package entry

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Image struct {
	File *File

	Date time.Time
	Dims *dims
	Info Info
}

func NewImage(path string, hold *Hold) (*Image, error) {
	path = strings.Replace(path, "cache/1600/", "", -1)

	file, err := NewFile(path, hold)
	if err != nil {
		return nil, err
	}

	dims, err := loadDims(path)
	if err != nil {
		return nil, err
	}

	info, err := loadFileInfo(path)
	if err != nil {
		return nil, err
	}

	date, err := loadImageDate(path)
	if err != nil {
		// if Debug {
		// 	log.Println(err)
		// }
		// return nil, err
	}

	return &Image{
		File: file,
		Date: date,
		Dims: dims,
		Info: info,
	}, nil
}

type dims struct {
	width, height int
}

func loadDims(path string) (*dims, error) {
	// /dir/file.jpg/cache/dims/file.jpg.txt
	b, err := ioutil.ReadFile(dimsFile(path))
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(b))
	x := strings.Index(s, "x")
	if x == -1 || len(s) < x+1 {
		return nil, fmt.Errorf("invalid dimensions %v", path)
	}

	w := s[:x]
	h := s[x+1:]

	width, err := strconv.Atoi(w)
	if err != nil {
		return nil, err
	}

	height, err := strconv.Atoi(h)
	if err != nil {
		return nil, err
	}

	return &dims{width: width, height: height}, nil
}

func dimsFile(path string) string {
	return filepath.Join(cacheFolder(path), "dims", filepath.Base(path)+".txt")
}

func cacheFolder(path string) string {
	return filepath.Join(filepath.Dir(path), "cache")
}

func loadImageDate(path string) (time.Time, error) {
	return getFilenameDate(path)
}
