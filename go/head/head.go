package head

import (
	"fmt"
	"stferal/go/entry"
)

type Head struct {
	Title   string
	Section string
	Path    string
	Host    string

	Entry   entry.Entry
	Options map[string]bool

	/*
		Local   bool
		Dark    bool
		Large   bool
		NoLog   bool
	*/

	Nav   Nav
	Lang  string
	Langs Langs

	Desc   string
	Schema *Schema
}

func (h *Head) Process() error {
	// TODO: check for nil entry?
	h.Lang = Lang(h.Host)
	h.Desc = h.GetDesc()
	h.Langs = h.MakeLangs()
	h.Nav = h.MakeNav()

	return nil
}

func SiteName(lang string) string {
	return "Sacer Feral"
}

func (h *Head) PageTitle() string {
	name := SiteName(h.Lang)
	if h.Title == "" {
		return name
	}
	return fmt.Sprintf("%v - %v", h.Title, name)
}

func (h *Head) PageURL() string {
	if l := h.Langs.Hreflang(h.Lang); l != nil {
		return l.Href
	}
	return ""
}

func (h *Head) DontIndex() bool {
	switch h.Path {
	case "/impressum", "/legal", "/privacy", "/datenschutz":
		return true
	}
	return false
}
