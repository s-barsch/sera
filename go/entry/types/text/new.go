package text

import (
	"time"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/file"
	"g.rg-s.com/sera/go/entry/info"
	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/tools/script"
)

type Text struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	raw script.LangMap

	Script *script.Script
}

func (t *Text) Copy() *Text {
	return &Text{
		parent: t.parent, // TODO: is this dangerous?
		file:   t.file,

		date: t.date,
		info: t.info.Copy(),

		raw: t.raw.Copy(),

		Script: t.Script.Copy(),
	}
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

	raw := langs.Copy()
	script := script.RenderScript(langs)

	return &Text{
		parent: parent,
		file:   file,

		date: date,
		info: inf,

		raw: raw,

		Script: script,
	}, nil
}

func ReadTextFile(path string) (info.Info, script.LangMap, error) {
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
