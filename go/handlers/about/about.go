package about

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type aboutTree struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func About(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	about := s.Store.Trees["about"].Access(m.Auth.Subscriber)[m.Lang]

	t, err := about.SearchTree(m.Split.Slug, m.Lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	m.Title = t.Title(m.Lang)
	m.SetSection("about")
	m.SetHreflang(t)

	err = s.Store.ExecuteTemplate(w, aboutTemplate(t.Level()), &aboutTree{
		Meta: m,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}

func aboutTemplate(level int) string {
	if level == 0 {
		return "about-main"
	}
	return "about-page"
}

func Rewrites(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	folder := m.Path[:len("/about")]
	if folder == "/about" {
		http.Redirect(w, r, "/en"+m.Path, http.StatusMovedPermanently)
		return
	}
}
