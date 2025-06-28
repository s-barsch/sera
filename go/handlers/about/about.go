package about

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

type aboutTree struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func About(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		about := v.Store.Trees["about"].Access(m.Auth.Subscriber)[m.Lang]

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

		err = v.Engine.ExecuteTemplate(w, aboutTemplate(t.Level()), &aboutTree{
			Meta: m,
			Tree: t,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func aboutTemplate(level int) string {
	if level == 0 {
		return "about-main"
	}
	return "about-page"
}

func Rewrites(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		folder := m.Path[:len("/about")]
		if folder == "/about" {
			http.Redirect(w, r, "/en"+m.Path, http.StatusMovedPermanently)
			return
		}
	}
}
