package about

import (
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	about := s.Trees["about"].Access(m.Auth.Subscriber)[m.Lang]

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
		http.Redirect(w, r, "/en"+m.Path, http.StatusMovedPermanently)
		return
	}
	/*
		if folder == "/ueber" {
			http.Redirect(w, r, "/de"+m.Path, http.StatusMovedPermanently)
			return
		}
	*/
}
