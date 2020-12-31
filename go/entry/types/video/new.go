package video

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sacer/go/entry"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/audio"
	"sacer/go/entry/types/text"
	"sacer/go/server/paths"
	"strconv"
	"strings"
	"time"
)

type Video struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Sources []*Source

	Subtitles  []string
	Transcript *text.Script
}

type Source struct {
	Path string
	Size string 
}

func NewVideo(path string, parent entry.Entry) (*Video, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "NewVideo",
	}

	file, err := file.NewFile(path)
	if err != nil {
		return nil, err
	}

	sources, err := getSources(path)
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

	script := audio.GetTranscript(inf)

	return &Video{
		parent:     parent,
		file:       file,
		date:       date,
		info:       inf,
		Subtitles:  subs,
		Transcript: script,
		Sources:    sources,
	}, nil
}

func getSubtitles(path string) []string {
	dir := filepath.Dir(path)
	name := stripSize(tools.StripExt(filepath.Base(path)))
	langs := []string{}
	for _, lang := range []string{"de", "en"} {
		_, err := os.Stat(filepath.Join(dir, "vtt", fmt.Sprintf("%v-%v.vtt", name, lang)))
		if err == nil {
			langs = append(langs, lang)
		}
	}
	return langs
}


func getSources(path string) ([]*Source, error) {

	s, _ := getSource(filepath.Base(path), true)

	sources := []*Source{s}

	sizes := filepath.Dir(path) + "/sizes"
	_, err := os.Stat(sizes)
	if err != nil {
		return sources, nil
	}

	l, err := ioutil.ReadDir(sizes)
	if err != nil {
		return nil, err
	}

	for _, fi := range l {
		if fi.IsDir() {
			continue
		}

		s, err := getSource(filepath.Join("sizes", fi.Name()), false)
		if err != nil {
			return nil, err
		}

		sources = append(sources, s)
	}
	return sources, nil
}

func getSource(path string, isTop bool) (*Source, error) {
	subfile := paths.SplitSubpath(path)
	size, err := strconv.Atoi(subfile.Size)
	if err != nil {
		if isTop {
			size = 1080
		} else {
			return nil, fmt.Errorf("getSources: Could not find size of %v", path)
		}
	}
	return &Source{
		Path: path,
		Size: strconv.Itoa(size),
	}, nil
}

func stripSize(name string) string {
	i := strings.LastIndex(name, "-")
	if i > 0 {
		return name[:i]
	}
	return name
}
