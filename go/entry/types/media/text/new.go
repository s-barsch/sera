package text

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

type Text struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	text  map[string]string
	blank map[string]string
}

func NewText(path string, parent entry.Entry) (*Text, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewText",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf, parts, err := readTextFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		date, err = helper.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	return &Text{
		parent: parent,
		file:   file,

		date: date,
		info: inf,

		text:  parts,
		blank: parts,
	}, nil
}

func readTextFile(path string) (info.Info, map[string]string, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readTextFile",
	}

	parts, err := splitTextFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, nil, fnErr
	}

	inf, err := info.UnmarshalInfo([]byte(parts["info"]))
	if err != nil {
		fnErr.Err = err
		return nil, nil, fnErr
	}

	delete(parts, "info")
	return inf, parts, nil
}
