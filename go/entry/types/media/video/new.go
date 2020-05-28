package video

import (
	"fmt"
	"os"
	"path/filepath"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"stferal/go/entry/parts/file"
	"time"
)

type Video struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Subtitles []string
}

func NewVideo(path string, parent entry.Entry) (*Video, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewVideo",
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

	date, err := helper.ParseTimestamp(inf["date"])
	if err != nil {
		date, err = helper.ParseDatePath(path)
		if err != nil {
			fnErr.Err = err
			return nil, fnErr
		}
	}

	subs := getSubtitles(path)

	return &Video{
		parent: parent,
		file:   file,
		date:   date,
		info:   inf,
		Subtitles: subs,
	}, nil
}

func getSubtitles(path string) []string {
	dir := filepath.Dir(path)
	name := helper.StripExt(filepath.Base(path))
	langs := []string{}
	for _, lang := range []string{"de", "en"} {
		_, err := os.Stat(filepath.Join(dir, "vtt", fmt.Sprintf("%v-%v.vtt", name, lang)))
		if err == nil {
			langs = append(langs, lang)
		}
	}
	return langs
}

