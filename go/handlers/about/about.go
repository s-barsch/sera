package about

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
)

type aboutTree struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func About(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	about := s.Trees["about"].Access(m.Auth.Subscriber)[m.Lang]

	p := paths.Split(m.Path)
	t, err := about.SearchTree(p.Slug, m.Lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	m.Title = t.Title(m.Lang)
	m.Section = "about"

	err = m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, aboutTemplate(t.Level()), &aboutTree{
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
