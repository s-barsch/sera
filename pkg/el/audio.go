package el

import (
	"fmt"
)

func (a *Audio) Acronym() string {
	return ToB16(a.Date)
}

func (a *Audio) AcronymShort() string {
	return shortenAcronym(a.Acronym())
}

func (a *Audio) Title(lang string) string {
	t := a.Info.Title(lang)
	if t == "" {
		return a.AcronymShort()
	}
	return t
}

func (a *Audio) Permalink(lang string) string {
	if a.Info.Title(lang) == "" {
		return fmt.Sprintf("%v/%v", a.File.Hold.Path(lang), ToB16(a.Date))
	}
	return fmt.Sprintf("%v/%v-%v", a.File.Hold.Path(lang), a.Info.Slug(lang), ToB16(a.Date))
}

func (v *Audio) CaptionPath(lang string) string {
	return fmt.Sprintf(
		"%v/files/vtt/%v-%v.vtt",
		v.File.Hold.Permalink(lang),
		v.File.BaseNoExt(),
		lang,
	)
}

func NewAudio(path string, hold *Hold) (*Audio, error) {
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
		//TODO: what is happening here?

		// if Debug {
		// 	log.Println(err)
		// }
		// return nil, err
	}

	subs := getSubtitles(path)

	return &Audio{
		File: file,
		Date: date,
		Info: info,

		Subtitles: subs,
	}, nil
}
