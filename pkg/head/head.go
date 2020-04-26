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
	Dark   bool
	Large   bool
	NoLog   bool
	El      interface{}

	Nav   Nav
	Lang  string
	Langs Langs

	Desc   string
	Schema *Schema
}

func (h *Head) TypeModeTitle(lang string) string {
	switch lang {
	case "en":
		if h.Dark {
			return "Switch to Default type mode"
		} else {
			return "Switch to Large type mode"
		}
	default:
		if h.Dark {
			return "Wechsle zum Großschrift-Modus"
		} else {
			return "Wechsle zum Standardschrift-Modus"
		}
	}
}

func (h *Head) DarkLinkTitle(lang string) string {
	switch lang {
	case "en":
		if h.Dark {
			return "Switch to Light mode"
		} else {
			return "Switch to Dark mode"
		}
	default:
		if h.Dark {
			return "Wechsle zu Hellmodus"
		} else {
			return "Wechsle zu Dunkelmodus"
		}
	}
}

func (h *Head) TypeModeLink(lang string) string {
	switch lang {
	case "en":
		if h.Large {
			return "/opt/defaulttype/"
		} else {
			return "/opt/largetype/"
		}
	default:
		if h.Large {
			return "/opt/standardschrift/"
		} else {
			return "/opt/grossschrift/"
		}
	}
}

func (h *Head) DarkLink(lang string) string {
	switch lang {
	case "en":
		if h.Dark {
			return "/opt/lightmode/"
		} else {
			return "/opt/darkmode/"
		}
	default:
		if h.Dark {
			return "/opt/hellmodus/"
		} else {
			return "/opt/dunkelmodus/"
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

func LogMode(r *http.Request) bool {
	c, err := r.Cookie("nolog")
	if err != nil {
		return false
	}
	return c.Value == "true"
}

func TypeMode(r *http.Request) bool {
	c, err := r.Cookie("largetype")
	if err != nil {
		return false
	}
	return c.Value == "true"
}

func DarkMode(r *http.Request) bool {
	c, err := r.Cookie("darkmode")
	if err != nil {
		return false
	}
	return c.Value == "true"
}
