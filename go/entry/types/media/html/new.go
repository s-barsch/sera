package html 

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"stferal/go/entry/types/media/text"
	"time"
)

type Html struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Html map[string]string
}

func NewHtml(path string, parent entry.Entry) (*Html, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewHtml",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf, langs, err := text.ReadTextFile(path)
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

	de := langs["de"]
	en := langs["en"]

	return &Html{
		parent: parent,
		file:   file,

		date: date,
		info: inf,

		Html: map[string]string{
			"de": de,
			"en": en,
		},
	}, nil
}


