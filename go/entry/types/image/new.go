package image

import (
	"fmt"
	p "path/filepath"
	"strings"
	"time"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/file"
	"g.sacerb.com/sacer/go/entry/info"
	"g.sacerb.com/sacer/go/entry/tools"
)

type Image struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Dims *Dims
}

func NewImage(path string, parent entry.Entry) (*Image, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "NewImage",
	}

	path = strings.Replace(path, "img/1600/", "", -1)

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	dims, err := loadDims(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf := info.Info{}

	if info.HasFileInfo(path) {
		i, err := info.ReadFileInfo(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
		inf = i
	}

	date, err := getImageDate(path, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return &Image{
		parent: parent,
		file:   file,
		date:   date,
		info:   inf,
		Dims:   dims,
	}, nil
}

func getImageDate(path string, parent entry.Entry) (time.Time, error) {
	if p.Base(path) == "cover.jpg" {
		if parent == nil {
			return time.Time{}, fmt.Errorf("getImageDate: parent shall not be nil")
		}
		return parent.Date().Add(time.Second), nil
	}
	return tools.ParseDatePath(path)
}

func (i *Image) Copy() *Image {
	return &Image{
		parent: i.parent,
		file:   i.file.Copy(),

		date: i.date,
		info: i.info.Copy(),

		Dims: i.Dims,
	}
}

func (i *Image) Blur() *Image {
	i = i.Copy()

	i.file.Path = addBlur(i.file.Path)

	return i
}

func addBlur(path string) string {
	i := strings.LastIndex(path, ".")
	if i < 0 {
		panic("addBlur: invalid path: " + path)
	}
	return path[:i] + "_blur" + path[i:]
}
