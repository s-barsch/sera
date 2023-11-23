package meta

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

func (m *Meta) MakeNav() Nav {
	nav := NewNav(m.Lang)

	section := sectionLang(m.Section, m.Lang)

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

func sectionLang(section, lang string) string {
	switch section {
	case "about":
		return tools.AboutName[lang]
	case "reels":
		return tools.KineName[lang]
	}
	return section
}

func NewNav(lang string) Nav {
	aboutName := tools.AboutName[lang]
	reelsName := tools.KineName[lang]
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
			Name: reelsName,
			Href: fmt.Sprintf("/%v/%v", lang, reelsName),
		},
		{
			Name: aboutName,
			Href: fmt.Sprintf("/%v/%v", lang, sanitize.Name(aboutName)),
		},
	}
}
