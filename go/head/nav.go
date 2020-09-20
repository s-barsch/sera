package head

import (
	"fmt"
	"stferal/go/entry/helper"

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
		return helper.AboutName[lang]
	case "kine":
		return helper.KineName[lang]
	}
	return section
}

func NewNav(lang string) Nav {
	about := helper.AboutName[lang]
	kine := helper.KineName[lang]
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
		&Link{
			Name: kine,
			Href: fmt.Sprintf("/%v", kine),
		},
		&Link{
			Name: about,
			Href: fmt.Sprintf("/%v", sanitize.Name(about)),
		},
	}
}
