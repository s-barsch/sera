package entry

import (
	"fmt"
	//"st/el/hyph"
)

type Html struct {
	File *File

	Date time.Time
	Info Info

	Html map[string]string
}

func (ht *Html) Acronym() string {
	return ToB16(ht.Date)
}

func (ht *Html) AcronymShort() string {
	return shortenAcronym(ht.Acronym())
}

func (ht *Html) Title(lang string) string {
	title := ht.Info.Title(lang)
	if title != "" {
		return title
	}
	return ht.AcronymShort()
}

func (ht *Html) Permalink(lang string) string {
	return fmt.Sprintf("%v/%v", ht.File.Hold.Path(lang), ToB16(ht.Date))
}

/* copied from NewFileText */
func NewHtml(path string, hold *Hold) (*Html, error) {
	file, err := NewFile(path, hold)
	if err != nil {
		return nil, err
	}

	parts, err := splitSingleText(path)
	if err != nil {
		return nil, err
	}

	info, err := unmarshalInfo([]byte(parts["info"]))
	if err != nil {
		return nil, fmt.Errorf("%v (%v)", err, path)
	}

	date, err := ParseDate(Shorten(helper.StripExt(file.Base())))
	if err != nil {
		date, err = ParseDate(info["date"])
		if err != nil {
			return nil, fmt.Errorf("Cannot read date of %v\nErr: %v", path, err)
		}
	}

	de := parts["de"]
	en := parts["en"]

	/*
		if info["hyphen"] == "true" {
			deH, err := HyphenateText([]byte(de), "de")
			if err != nil {
				return nil, err
			}
			de = string(deH)

			enH, err := HyphenateText([]byte(en), "en")
			if err != nil {
				return nil, err
			}
			en = string(enH)
		}
	*/

	file.Id = date.Format(Timestamp)

	return &Html{
		File: file,
		Info: info,
		Date: date,
		Html: map[string]string{
			"de": de,
			"en": en,
		},
	}, nil
}
