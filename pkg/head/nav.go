package head

import (
	"fmt"
	"github.com/kennygrant/sanitize"
)

type Nav []*Link

type Link struct {
	IsActive   bool
	Name, Href string
}

func (h *Head) MakeNav() Nav {
	nav := NewNav(h.Lang)
	section := h.LocalSection(h.Section)

	for _, l := range nav {
		if section == l.Name {
			l.IsActive = true
		}
	}

	return nav
}

var aboutNames = map[string]string{
	"de": "über",
	"en": "about",
}

func (h *Head) LocalSection(name string) string {
	if h.Lang == "de" && h.Section == "about" {
		return "über"
	}
	return name
}

func NewNav(lang string) Nav {
	about := aboutNames[lang]
	return []*Link{
		&Link{
			Name: "home",
			Href: "/",
		},
		&Link{
			Name: "index",
			Href: "/index/",
		},
		&Link{
			Name: "graph",
			Href: "/graph/",
		},
		&Link{
			Name: about,
			Href: fmt.Sprintf("/%v/", sanitize.Name(about)),
		},
	}
}
