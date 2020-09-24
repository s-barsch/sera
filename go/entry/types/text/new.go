package text

import (
	"sacer/go/entry"
	"sacer/go/entry/helper"
	"sacer/go/entry/helper/markup"
	"sacer/go/entry/parts/file"
	"sacer/go/entry/parts/info"
	"time"
)

type Text struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	TextLangs map[string]string
	blank     map[string]string

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

	inf, langs, err := ReadTextFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseTimestamp(inf["date"])
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
		blank:     langs,

		Notes: notes,
	}, nil
}


func renderLangs(langs map[string]string) (map[string]string, map[string][]string) {
	notes := map[string][]string{}
	for _, l := range []string{"de", "en"} {
		text, ns := markup.Render(langs[l])
		text = helper.RenderMarkdown(text)
		langs[l] = text
		notes[l] = ns
	}
	return langs, notes
}

func ReadTextFile(path string) (info.Info, map[string]string, error) {
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
