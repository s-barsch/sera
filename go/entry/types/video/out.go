package video

import (
	"fmt"
	"path/filepath"
)

func (v *Video) Location(size string) (string, error) {
	if size == "" {
		size = "1080"
	}
	for _, s := range v.Sources {
		if s.Size == size {
			return filepath.Join(v.file.Dir(), s.Path), nil
		}
	}
	return "", fmt.Errorf("Cannot find size %v in %v", size, v.file.Path)
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

func (v *Video) SubtitlePath(lang string) string {
	parent := v.Perma(lang)
	if v.parent.Type() == "set" {
		parent = v.parent.Perma(lang)
	}
	return vttPath(parent + "/files", stripSize(v.file.NameNoExt()), lang)
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

func (v *Video) SubtitleLocation(lang string) string {
	return vttPath(v.file.Dir(), v.file.NameNoExt(), lang)
}

func hlsPath(path, nameNoExt string) string {
	return fmt.Sprintf("%v/hls/%v.m3u8", path, nameNoExt)
}

func vttPath(path, nameNoExt, lang string) string {
	return fmt.Sprintf("%v/vtt/%v-%v.vtt", path, nameNoExt, lang)
}

