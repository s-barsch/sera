package about

import (
	"net/http"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
	"sacer/go/server"
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
