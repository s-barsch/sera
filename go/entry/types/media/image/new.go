package image

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"stferal/go/entry/parts/file"
	"strings"
	"time"
)

type Image struct {
	File *file.File

	Date time.Time
	// Dims *dims
	Info info.Info
}

func NewImage(path string) (*Image, error) {
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

	date, err := helper.ParseDatePath(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return &Image{
		File: file,
		Date: date,
		//Dims: dims,
		Info: inf,
	}, nil
}

