package info

import (
	"fmt"
	"sacer/go/entry/tools"
	"time"
)

func (i Info) Title(lang string) string {
	return i.Field("title", lang)
}

func (i Info) HyphTitle(lang string) string {
	return i.Field("title-hyph", lang)
}

func (i Info) Blur() bool {
	return i["blur"] == "true"
}

func (i Info) Private() bool {
	return i["private"] == "true"
}

func (i Info) Note() bool {
	return i["note"] == "true"
}

func (i Info) Wall() bool {
	return i["wall"] == "true"
}

func (i Info) Caption(lang string) string {
	return i.Field("caption", lang)
}

func (i Info) Description(lang string) string {
	return i.Field("description", lang)
}

func (i Info) Alt(lang string) string {
	return i.Field("alt", lang)
}

func (i Info) Slug(lang string) string {
	return i.Field("slug", lang)
}

func (i Info) TextStyle() string {
	if s := i["style"]; s != "" {
		return s
	}
	return "indent"
}

func (i Info) Field(key, lang string) string {
	if lang != "de" {
		return i[key+"-"+lang]
	}
	return i[key]
}

type DateWrapper struct {
	Date time.Time
}

func (d *DateWrapper) Format(format string) string {
	if d == nil {
		return ""
	}
	return d.Date.Format(format)
}

func (i Info) Published() *DateWrapper {
	p := i["published"]
	if p == "" {
		return nil
	}
	return parseDate(p)
}

func (i Info) Revision(format string) *DateWrapper {
	r := i["revision"]
	if r == "" {
		return nil
	}
	return parseDate(r)
}

func parseDate(field string) *DateWrapper {
	d, err := time.Parse(tools.Timestamp, field)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &DateWrapper{
		Date: d,
	}
}

/*
func (i Info) Label(lang string) string {
	if label := i.Field("label", lang); label != "" {
		return label
	}
	return i.Title(lang)
}

func (i Info) TitleUpper(lang string) string {
	title := i.Field("title-hyph", lang)
	return s.Replace(title, "ß", "ss", -1)
}
*/
