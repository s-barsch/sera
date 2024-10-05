package meta

import (
	"fmt"
)

type Nav []*Link

type Link struct {
	IsActive   bool
	Name, Href string
}

func (m *Meta) MakeNav(section string) Nav {
	nav := NewNav(m.Lang)

	for _, l := range nav {
		if section == l.Name {
			l.IsActive = true
		}
		if section == "komposita" && l.Name == "index" {
			l.IsActive = true
		}
	}

	return nav
}

func NewNav(lang string) Nav {
	homeLang := ""
	if lang == "de" {
		homeLang = "de"
	}
	return []*Link{
		{
			Name: "home",
			Href: fmt.Sprintf("/%v", homeLang),
		},
		{
			Name: "graph",
			Href: fmt.Sprintf("/%v/graph", lang),
		},
		{
			Name: "cache",
			Href: fmt.Sprintf("/%v/%v", lang, "cache"),
		},
		{
			Name: "about",
			Href: fmt.Sprintf("/%v/%v", lang, "about"),
		},
	}
}
