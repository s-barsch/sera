package html

import (
	"time"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/file"
	"g.sacerb.com/sacer/go/entry/info"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/entry/types/text"
)

type Html struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Html map[string]string
}

func NewHtml(path string, parent entry.Entry) (*Html, error) {
	fnErr := &tools.Err{
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

	date, err := tools.ParseTimestamp(inf["date"])
	if err != nil {
		date, err = tools.ParseDatePath(path)
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
