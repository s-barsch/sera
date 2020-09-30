package text

import (
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/markup"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
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
	fnErr := &tools.Err{
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

	date, err := tools.ParseTimestamp(inf["date"])
	if err != nil {
		date, err = tools.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	rendered, notes := markupLangs(langs)

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


func markupLangs(langs map[string]string) (map[string]string, map[string][]string) {
	notes := map[string][]string{}
	for lang, _ := range tools.Langs {
		text, ns := markup.Render(langs[lang])
		text = tools.RenderMarkdown(text)
		langs[lang] = text
		notes[lang] = ns
	}
	return langs, notes
}

func ReadTextFile(path string) (info.Info, map[string]string, error) {
	fnErr := &tools.Err{
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
