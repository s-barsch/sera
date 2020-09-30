package text

import (
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/hyph"
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

	Langs map[string]string
	Notes map[string][]string
}

type Langs map[string]string

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

	notes := langs.Markup()

	langs.Markdown()
	langs.Hyphenate()

	return &Text{
		parent: parent,
		file:   file,

		date: date,
		info: inf,

		Langs: langs,

		Notes: renderNotes(notes),
	}, nil
}

func renderNotes(notes map[string][]string) map[string][]string {
	for l, _ := range tools.Langs {
		for i, _ := range notes[l] {
			notes[l][i] = tools.Markdown(notes[l][i])
			notes[l][i] = hyph.Hyphenate(notes[l][i], l)
		}
	}
	return notes
}

func (langs Langs) Hyphenate() {
	for l, _ := range tools.Langs {
		langs[l] = hyph.Hyphenate(langs[l], l)
	}
}

func (langs Langs) Markdown() {
	for l, _ := range tools.Langs {
		langs[l] = tools.Markdown(langs[l])
	}
}

func (langs Langs) Markup() map[string][]string {
	notes := map[string][]string{}

	for l, _ := range tools.Langs {
		text, ns := markup.Render(langs[l])
		langs[l] = text
		notes[l] = ns
	}

	return notes
}

func ReadTextFile(path string) (info.Info, Langs, error) {
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
