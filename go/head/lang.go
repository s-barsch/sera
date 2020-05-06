package head

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/types/media/text"
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
	for _, lang := range []string{"de", "en"} {
		langs = append(langs, getLink(h, h.Entry, lang))
	}
	return langs
}

func getLink(h *Head, e entry.Entry, lang string) *Link {
	href := h.AbsoluteURL(e.Perma(lang), lang)

	if lang != "de" && !isTranslated(e, lang) {
		href = h.HostAddress(lang)
	}

	return &Link{
		Name: lang,
		Href: href,
	}
}

func isTranslated(e entry.Entry, lang string) bool {
	if e.Info()["translated"] == "false" {
		return false
	}
	if txt, ok := e.(*text.Text); ok && txt.Text(lang) == "" {
		return false
	}
	return true
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
