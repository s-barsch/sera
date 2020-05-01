package text

/*
import (
	"fmt"
)

func (t *Text) Acronym() string {
	return ToB16(t.Date)
}

func (t *Text) AcronymShort() string {
	return shortenAcronym(t.Acronym())
}

func (t *Text) Title(lang string) string {
	title := t.Info.Title(lang)
	if title != "" {
		return title
	}
	return t.AcronymShort()
}

func (t *Text) TextLang(lang string) string {
	return t.Text[lang]
}

func (t *Text) Permalink(lang string) string {
	switch t.File.Section() {
	case "about":
		return t.File.Hold.Permalink(lang)
	case "index":
		return fmt.Sprintf("%v#%v", t.File.Hold.Permalink(lang), Normalize(t.Title(lang)))
	case "graph":
		if t.Info.Title(lang) == "" {
			return fmt.Sprintf("%v/%v", t.File.Hold.Path(lang), ToB16(t.Date))
		}
	}
	return fmt.Sprintf("%v/%v-%v", t.File.Hold.Path(lang), t.Info.Slug(lang), ToB16(t.Date))
}

func (t *Text) MetaDesc(lang string) string {
	return shortenDesc(t.Blank[lang], 300)
}

func shortenDesc(str string, length int) string {
	desc := ""
	linebreak := false

	for i, r := range str {
		if i >= length && (r == ' ' || r == ',') {
			return desc + "…"
		}

		if r == '\n' {
			if !linebreak {
				desc += " "
			}
			linebreak = true
			continue
		}

		linebreak = false
		desc += string(r)
	}
	return desc
}
*/
