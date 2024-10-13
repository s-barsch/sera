package video

import (
	"fmt"
	"math"
	"path/filepath"

	"g.rg-s.com/sera/go/entry/tools"
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
	return math.Round(mb*10) / 10
}

// ext = mp4 | vtt
// option = 1280, 720 | de, en
func (v *Video) Location(ext, opt string) (string, error) {
	if ext == "vtt" {
		return v.CaptionLocation(opt), nil
	}
	if opt == "err" {
		return "", fmt.Errorf("(*Video) Location(): faulty input: %v", v.File().Path)
	}
	if opt == "" {
		opt = "1080"
	}
	for _, s := range v.Sources {
		if s.Resolution == opt {
			return filepath.Join(v.file.Dir(), s.Path), nil
		}
	}
	return "", fmt.Errorf("cannot find resolution %v in %v", opt, v.file.Path)
}

func (v *Video) CaptionLocation(lang string) string {
	return tools.VTTPath(v.file.Dir(), v.file.NameNoExt(), lang)
}

func (v *Video) FilesPath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return fmt.Sprintf("%v/files", parent)
}

func (v *Video) CaptionPath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return tools.VTTPath(parent+"/files", stripResolution(v.file.NameNoExt()), lang)
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

func (v *Video) Captioned() bool {
	return len(v.Captions) == 2
}

func (v *Video) Transcripted() bool {
	for _, str := range v.Transcript.LangMap {
		if str == "" {
			return false
		}
	}
	return true
}
