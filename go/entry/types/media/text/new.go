package text

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

type Text struct {
	File *file.File

	Date time.Time
	Info info.Info

	Text  map[string]string
	Blank map[string]string
}

func NewText(path string) (*Text, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewText",
	}

	file, err := file.New(path)
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
		File: file,

		Date: date,
		Info: inf,

		Text:  parts,
		Blank: parts,
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
