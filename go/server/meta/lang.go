package meta

import (
	"fmt"
	"log"

	"g.rg-s.com/sera/go/entry"
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

func (langs Langs) Active(lang string) *Link {
	for _, l := range langs {
		if l.Name == lang {
			return l
		}
	}
	log.Printf("cannot find active link of lang: %v", lang)
	return nil
}

func (langs Langs) Sort(lang string) Langs {
	if langs[1].Name == lang {
		return langs
	}
	return Langs{langs[1], langs[0]}
}

func (langs Langs) Hreflang(name string) *Link {
	for _, l := range langs {
		if l.Name == name {
			return l
		}
	}
	return nil
}

func MakeHreflangs(host string, e entry.Entry) Langs {
	langs := []*Link{}
	for _, lang := range []string{"de", "en"} {
		langs = append(langs, getLink(host, e, lang))
	}
	return langs
}

func getLink(host string, e entry.Entry, lang string) *Link {
	path := ""
	if e == nil {
		path = homePath[lang]
	} else {
		path = e.Perma(lang)
	}

	return &Link{
		Name: lang,
		Href: fmt.Sprintf("%v%v", host, path),
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
		return "http://sera"
	}
	return "https://seraferal.com"
}

func (m *Meta) IsHostnameLocal() bool {
	return isHostnameLocal(m.Host)
}

func isHostnameLocal(host string) bool {
	switch host {
	case "localhost:8013", "sera":
		return true
	}
	return false
}

var homePath = map[string]string{
	"de": "/de",
	"en": "/",
}
