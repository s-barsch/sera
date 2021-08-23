package head

import (
	"fmt"
	"sacer/go/entry/tools"

	"github.com/kennygrant/sanitize"
)

type Nav []*Link

type Link struct {
	IsActive   bool
	Name, Href string
}

func (h *Head) MakeNav() Nav {
	nav := NewNav(h.Lang)

	section := sectionLang(h.Section, h.Lang)

	for _, l := range nav {
		if section == l.Name {
			l.IsActive = true
		}
	}

	return nav
}

func sectionLang(section, lang string) string {
	switch section {
	case "about":
		return tools.AboutName[lang]
	case "kine":
		return tools.KineName[lang]
	}
	return section
}

func NewNav(lang string) Nav {
	about := tools.AboutName[lang]
	//kine := tools.KineName[lang]
	return []*Link{
		&Link{
			Name: "home",
			Href: "/",
		},
		&Link{
			Name: "index",
			Href: "/index",
		},
		&Link{
			Name: "graph",
			Href: "/graph",
		},
		/*
		&Link{
			Name: kine,
			Href: fmt.Sprintf("/%v", kine),
		},
		*/
		&Link{
			Name: about,
			Href: fmt.Sprintf("/%v", sanitize.Name(about)),
		},
	}
}
