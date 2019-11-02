package head

import (
	"fmt"
	"net/http"
)

type Head struct {
	Title   string
	Section string
	Path    string
	Host    string
	Local   bool
	Night   bool
	El      interface{}

	Nav   Nav
	Lang  string
	Langs Langs

	Desc   string
	Schema *Schema
}

func (h *Head) NightLinkTitle(lang string) string {
	switch lang {
	case "en":
		if h.Night {
			return "Switch to Day mode"
		} else {
			return "Switch to Night mode"
		}
	default:
		if h.Night {
			return "Wechsle zu Tagmodus"
		} else {
			return "Wechsle zu Nachtmodus"
		}
	}
}

func (h *Head) NightLink(lang string) string {
	switch lang {
	case "en":
		if h.Night {
			return "/daymode/"
		} else {
			return "/nightmode/"
		}
	default:
		if h.Night {
			return "/tagmodus/"
		} else {
			return "/nachtmodus/"
		}
	}
}

func (h *Head) Make() error {
	// can check for missing entries?
	h.Lang = Lang(h.Host)
	h.Desc = h.GetDesc()
	h.Langs = h.MakeLangs()
	h.Nav = h.MakeNav()

	return nil
}

func (h *Head) PageTitle() string {
	if h.Title == "" {
		return "Stef Feral"
	}
	switch h.Lang {
	case "en":
		return fmt.Sprintf("%v - Stef Feral - English", h.Title)
	default:
		return fmt.Sprintf("%v - Stef Feral", h.Title)
	}
}

func (h *Head) PageURL() string {
	if l := h.Langs.Hreflang(h.Lang); l != nil {
		return l.Href
	}
	return ""
}

func (h *Head) DontIndex() bool {
	switch h.Path {
	case "/impressum/", "/legal/", "/privacy/", "/datenschutz/":
		return true
	}
	return false
}

func NightMode(r *http.Request) bool {
	c, err := r.Cookie("nightmode")
	if err != nil {
		return false
	}
	return c.Value == "true"
}
