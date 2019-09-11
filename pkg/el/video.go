package el

import (
	"fmt"
	"os"
	"path/filepath"
)

func (v *Video) Acronym() string {
	return ToB16(v.Date)
}

func (v *Video) AcronymShort() string {
	return shortenAcronym(v.Acronym())
}

func (v *Video) Title(lang string) string {
	t := v.Info.Title(lang)
	if t == "" {
		return v.AcronymShort()
	}
	return t
}

func (v *Video) Permalink(lang string) string {
	if v.Info.Title(lang) == "" {
		return fmt.Sprintf("%v/%v", v.File.Hold.Path(lang), ToB16(v.Date))
	}
	return fmt.Sprintf("%v/%v-%v", v.File.Hold.Path(lang), v.Info.Slug(lang), ToB16(v.Date))
}

func (v *Video) SubtitlePath(lang string) string {
	return fmt.Sprintf(
		"%v/files/vtt/%v-%v.vtt",
		v.File.Hold.Permalink(lang),
		v.File.BaseNoExt(),
		lang,
	)
}

func NewVideo(path string, hold *Hold) (*Video, error) {
	file, err := NewFile(path, hold)
	if err != nil {
		return nil, err
	}

	info, err := loadFileInfo(path)
	if err != nil {
		return nil, err
	}

	date, err := loadImageDate(path)
	if err != nil {
		// if Debug {
		// 	log.Println(err)
		// }
		// return nil, err
	}

	subs := getSubtitles(path)

	return &Video{
		File: file,
		Date: date,
		Info: info,

		Subtitles: subs,
	}, nil
}

func getSubtitles(path string) []string {
	dir := filepath.Dir(path)
	name := stripExt(filepath.Base(path))
	langs := []string{}
	for _, lang := range []string{"de", "en"} {
		_, err := os.Stat(filepath.Join(dir, "vtt", fmt.Sprintf("%v-%v.vtt", name, lang)))
		if err == nil {
			langs = append(langs, lang)
		}
	}
	return langs
}
