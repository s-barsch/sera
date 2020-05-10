package text

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/helper/markup"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	bf "gopkg.in/russross/blackfriday.v2"
	"time"
)

type Text struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	TextLangs map[string]string
	blank map[string]string

	Notes map[string][]string
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

	inf, langs, err := readTextFile(path)
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

	rendered, notes := renderLangs(langs)

	return &Text{
		parent: parent,
		file:   file,

		date: date,
		info: inf,

		TextLangs: rendered,
		blank: langs,

		Notes: notes,
	}, nil
}

func renderLangs(langs map[string]string) (map[string]string, map[string][]string) {
	notes := map[string][]string{}
	for _, l := range []string{"de", "en"} {
		text, ns := markup.Render(langs[l])
		text = string(bf.Run([]byte(text),bf.WithExtensions(bf.HardLineBreak)))
		langs[l] = text
		notes[l] = ns
	}
	return langs, notes
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
