package audio

import (
	"fmt"
	"os"
	p "path/filepath"
	"time"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/file"
	"g.sacerb.com/sacer/go/entry/info"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/entry/types/text"
)

type Audio struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Captions []string

	Transcript *text.Script
}

func NewAudio(path string, parent entry.Entry) (*Audio, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "NewAudio",
	}

	file, err := file.NewFile(path)
	if err != nil {
		return nil, err
	}

	inf := info.Info{}

	if info.HasFileInfo(path) {
		i, err := info.ReadFileInfo(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
		inf = i
	}

	date, err := tools.ParseTimestamp(inf["date"])
	if err != nil {
		date, err = tools.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	captions := getCaptions(path)

	script := GetTranscript(inf)

	return &Audio{
		parent:     parent,
		file:       file,
		date:       date,
		info:       inf,
		Captions:   captions,
		Transcript: script,
	}, nil
}

func getCaptions(path string) []string {
	dir := p.Dir(path)
	name := tools.StripExt(p.Base(path))
	langs := []string{}
	for _, lang := range []string{"de", "en"} {
		_, err := os.Stat(p.Join(dir, "vtt", fmt.Sprintf("%v-%v.vtt", name, lang)))
		if err == nil {
			langs = append(langs, lang)
		}
	}
	return langs
}

func GetTranscript(i info.Info) *text.Script {
	script := text.RenderScript(extractTranscript(i))
	script.NumberFootnotes(1)

	return script
}

func extractTranscript(i info.Info) text.Langs {
	langs := text.Langs{}
	for l := range tools.Langs {
		key := "transcript"
		if l != "de" {
			key += "-" + l
		}
		langs[l] = i[key]
	}
	return langs
}
