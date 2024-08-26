package about

import (
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	rel := m.Path[len("/about"):] // same length as "ueber"
	about := s.Trees["about"].Access(m.Auth.Subscriber)[m.Lang]

	if rel == "" {
		ServeAbout(s, w, r, m, about)
		return
	}

	p := paths.Split(m.Path)
	t, err := about.SearchTree(p.Slug, m.Lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ServeAbout(s, w, r, m, t)
}

func Rewrites(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	folder := m.Path[:len("/about")]
	if folder == "/about" {
		http.Redirect(w, r, "/en"+m.Path, 301)
		return
	}
	if folder == "/ueber" {
		http.Redirect(w, r, "/de"+m.Path, 301)
		return
	}
}
