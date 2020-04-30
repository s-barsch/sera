package text 

import (
	"fmt"
	"stferal/go/entry/file"
	"stferal/go/entry/info"
	"stferal/go/entry/helper"
	"time"
)

type Text struct {
	File *file.File

	Date time.Time
	Info info.Info

	Text  map[string]string
	Blank map[string]string
}

func New(path string) (*Text, error) {
	file, err := file.New(path)
	if err != nil {
		return nil, err
	}

	inf, parts, err := readTextFile(path)
	if err != nil {
		return nil, err
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		date, err = helper.ParseDatePath(path)
		if err != nil {
			return nil, err
		}
	}

	return &Text{
		File: file,

		Date: date,
		Info: inf,

		Text: parts,
		Blank: parts,
	}, nil
}

func readTextFile(path string) (inf info.Info, parts map[string]string, err error) {
	parts, err = splitTextFile(path)
	if err != nil {
		return
	}

	inf, err = info.UnmarshalInfo([]byte(parts["info"]))
	if err != nil {
		err = fmt.Errorf("%v (%v)", err, path)
		return
	}

	delete(parts, "info")
	return
}


