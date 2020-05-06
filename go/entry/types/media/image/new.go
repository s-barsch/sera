package image

import (
	"fmt"
	p "path/filepath"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"strings"
	"time"
)

type Image struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	// Dims *dims
}

func NewImage(path string, parent entry.Entry) (*Image, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewImage",
	}

	path = strings.Replace(path, "cache/1600/", "", -1)

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	/*
		dims, err := loadDims(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	*/

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
		//Dims: dims,
	}, nil
}

func getImageDate(path string, parent entry.Entry) (time.Time, error) {
	if p.Base(path) == "cover.jpg" {
		if parent == nil {
			return time.Time{}, fmt.Errorf("getImageDate: parent shall not be nil.")
		}
		return parent.Date().Add(time.Second), nil
	}
	return helper.ParseDatePath(path)
}
