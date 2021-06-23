package video

import (
	"fmt"
	"path/filepath"
	"math"
	"sacer/go/entry/tools"
)

func (v *Video) Ideal(res string) float64 {
	rate := float64(0)
	switch res {
	case "1080":
		rate = 1.2
	case "720":
		rate = 0.533
	case "480":
		rate = 0.075
	}
	mb := v.Duration * rate
	return math.Round(mb*10)/10
}

func (v *Video) Location(res string) (string, error) {
	if res == "" {
		res = "1080"
	}
	for _, s := range v.Sources {
		if s.Resolution == res {
			return filepath.Join(v.file.Dir(), s.Path), nil
		}
	}
	return "", fmt.Errorf("Cannot find resolution %v in %v", res, v.file.Path)
}

func (v *Video) FilesPath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return fmt.Sprintf("%v/files", parent)
}

/*
func (v *Video) FilePath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return fmt.Sprintf("%v/files/%v", parent, v.file.Name())
}
*/

func (v *Video) CaptionPath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return tools.VTTPath(parent + "/files", stripResolution(v.file.NameNoExt()), lang)
}

func (v *Video) CaptionsOn(captionsLang, pageLang string) bool {
	if pageLang != "de" {
		if captionsLang != "de" {
			return true
		}
	}
	return captionsLang == "de" && pageLang == "de" && v.Info()["captions-on"] == "true"
}

func (v *Video) HasCaptions(lang string) bool {
	for _, captionsLang := range v.Captions {
		if lang == captionsLang {
			return true
		}
	}
	return false
}

func (v *Video) CaptionLocation(lang string) string {
	return tools.VTTPath(v.file.Dir(), v.file.NameNoExt(), lang)
}

func hlsPath(path, nameNoExt string) string {
	return fmt.Sprintf("%v/hls/%v.m3u8", path, nameNoExt)
}
