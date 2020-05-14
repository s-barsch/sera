package audio

import (
	"os"
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
	p "path/filepath"
)

type Audio struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Subtitles []string
}

func NewAudio(path string, parent entry.Entry) (*Audio, error) {
	fnErr := &helper.Err{
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

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		date, err = helper.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	subs := getSubtitles(path)

	return &Audio{
		parent: parent,
		file:   file,
		date:   date,
		info:   inf,
		Subtitles: subs,
	}, nil
}

func getSubtitles(path string) []string {
	dir := p.Dir(path)
	name := helper.StripExt(p.Base(path))
	langs := []string{}
	for _, lang := range []string{"de", "en"} {
		_, err := os.Stat(p.Join(dir, "vtt", fmt.Sprintf("%v-%v.vtt", name, lang)))
		if err == nil {
			langs = append(langs, lang)
		}
	}
	return langs
}

