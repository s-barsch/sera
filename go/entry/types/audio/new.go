package audio

import (
	"fmt"
	"os"
	p "path/filepath"
	"sacer/go/entry"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/text"
	"time"
)

type Audio struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Subtitles []string

	Transcript *Transcript
}

type Transcript struct {
	Langs text.Langs
	Notes text.Notes
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

	subs := getSubtitles(path)

	transcript := GetTranscript(inf)

	return &Audio{
		parent:     parent,
		file:       file,
		date:       date,
		info:       inf,
		Subtitles:  subs,
		Transcript: transcript,
	}, nil
}

func getSubtitles(path string) []string {
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

func GetTranscript(i info.Info) *Transcript {
	if i["transcript"] == "" {
		return nil
	}

	langs, notes := text.RenderLangs(extractTranscript(i))

	return &Transcript{
		Langs: langs,
		Notes: notes,
	}
}

func extractTranscript(i info.Info) text.Langs {
	langs := text.Langs{}
	for l, _ := range tools.Langs {
		key := "transcript"
		if l != "de" {
			key += "-" + l
		}
		langs[l] = i[key]
	}
	return langs
}
