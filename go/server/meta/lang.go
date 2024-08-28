package meta

import (
	"fmt"

	"g.sacerb.com/sacer/go/entry"
)

func Lang(path string) string {
	if len(path) < 3 {
		return "en"
	}
	if path[:3] == "/de" {
		return "de"
	}
	return "en"
}

type Langs []*Link

func (langs Langs) Hreflang(name string) *Link {
	for _, l := range langs {
		if l.Name == name {
			return l
		}
	}
	return nil
}

func (m *Meta) MakeLangs(e entry.Entry) Langs {
	langs := []*Link{}
	for _, lang := range []string{"de", "en"} {
		langs = append(langs, getLink(m, e, lang))
	}
	return langs
}

func getLink(m *Meta, e entry.Entry, lang string) *Link {
	href := ""

	if e == nil {
		href = fmt.Sprintf("%v/%v", m.HostAddress(), homeAddress[lang])
	} else {
		href = m.AbsoluteURL(e.Perma(lang), lang)
	}

	return &Link{
		Name: lang,
		Href: href,
	}
}

/*
func isTranslated(e entry.Entry, lang string) bool {
	if e.Info()["translated"] == "false" {
		return false
	}
	if txt, ok := e.(*text.Text); ok && txt.Script.Langs[lang] == "" {
		return false
	}
	return true
}
*/

func (m *Meta) AbsoluteURL(path, lang string) string {
	return fmt.Sprintf("%v%v", m.HostAddress(), path)
}

func (m *Meta) HostAddress() string {
	if isHostnameLocal(m.Host) {
		return "http://sacer"
	}
	return "https://sacer.site"
}

func (m *Meta) IsHostnameLocal() bool {
	return isHostnameLocal(m.Host)
}

func isHostnameLocal(host string) bool {
	switch host {
	case "localhost:8013", "sacer":
		return true
	}
	return false
}

var homeAddress = map[string]string{
	"de": "de",
	"en": "",
}
