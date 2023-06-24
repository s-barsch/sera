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

func (i Info) Published(format string) string {
	p := i["published"]
	if p == "" {
		return ""
	}
	d, err := parseDate(p)
	if err != nil {
		return fmt.Errorf("Could’t parse published date: %v", err).Error()
	}
	return d.Format(format)
}

func (i Info) Revision(format string) string {
	r := i["revision"]
	if r == "" {
		return ""
	}
	d, err := parseDate(r)
	if err != nil {
		return fmt.Errorf("Could’t parse revision date: %v", err).Error()
	}
	return d.Format(format)
}

func parseDate(field string) (time.Time, error) {
	return time.Parse(tools.Timestamp, field)
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
