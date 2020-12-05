package video

import (
	"fmt"
)

func (v *Video) Location(arg string) string {
	return v.file.Path
}

func (v *Video) FilePath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return fmt.Sprintf("%v/files/%v", parent, v.file.Name())
}

func (v *Video) HLSPath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return hlsPath(parent + "/files", v.file.NameNoExt())
}

func (v *Video) SubtitlePath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return vttPath(parent + "/files", v.file.NameNoExt(), lang)
}

func (v *Video) SubtitlesOn(subLang, pageLang string) bool {
	if pageLang != "de" {
		if subLang != "de" {
			return true
		}
	}
	return subLang == "de" && pageLang == "de" && v.Info()["subtitles-on"] == "true"
}

func (v *Video) HasSubtitles(lang string) bool {
	for _, subLang := range v.Subtitles {
		if lang == subLang {
			return true
		}
	}
	return false
}

func (v *Video) HLSLocation() string {
	return hlsPath(v.file.Dir(), v.file.NameNoExt())
}

func (v *Video) SubtitleLocation(lang string) string {
	return vttPath(v.file.Dir(), v.file.NameNoExt(), lang)
}

func hlsPath(path, nameNoExt string) string {
	return fmt.Sprintf("%v/hls/%v.m3u8", path, nameNoExt)
}

func vttPath(path, nameNoExt, lang string) string {
	return fmt.Sprintf("%v/vtt/%v-%v.vtt", path, nameNoExt, lang)
}

