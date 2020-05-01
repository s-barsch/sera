package head

import (
	"fmt"
	//"stferal/go/entry"
	//"stferal/go/entry/resolver"
)

func Lang(host string) string {
	switch host {
	case "en.st", "en.stferal.com":
		return "en"
	default:
		return "de"
	}
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

func (h *Head) MakeLangs() Langs {
	langs := []*Link{}
	/*
	for _, lang := range []string{"de", "en"} {
		perma, err := resolver.Perma(h.Entry, lang)
		panic(err)
		href := h.AbsoluteURL(perma, lang)
		if lang != "de" {
			if entry, ok := h.Entry.(*entry.Text); ok && entry.Text[lang] == "" {
				href = h.HostAddress(lang)
			}
			if hold, ok := h.Entry.(*entry.Hold); ok && hold.Info["translated"] == "false" {
				href = h.HostAddress(lang)
			}
		}
		langs = append(langs, &Link{
			Name: lang,
			Href: href,
		})
	}
	*/
	return langs
}

func (h *Head) AbsoluteURL(path, lang string) string {
	return fmt.Sprintf("%v%v", h.HostAddress(lang), path)
}

func (h *Head) HostAddress(lang string) string {
	if isLocal(h.Host) {
		return fmt.Sprintf("http://%v", hostsLocal[lang])
	}
	return fmt.Sprintf("https://%v", hosts[lang])
}

func isLocal(host string) bool {
	switch host {
	case "st", "en.st":
		return true
	}
	return false
}

var hosts = map[string]string{
	"de": "stferal.com",
	"en": "en.stferal.com",
}

var hostsLocal = map[string]string{
	"de": "st",
	"en": "en.st",
}
