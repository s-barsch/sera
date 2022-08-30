package about

import (
	"net/http"
	//"sacer/go/entry"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"sacer/go/server"
	"sacer/go/server/users"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := p[len("/about"):] // same length as "ueber"
	lang := head.Lang(r.Host)
	about := s.Trees["about"].Access(a.Subscriber)[lang]

	if rel == "" {
		ServeAbout(s, w, r, a, about)
		return
	}

	path := paths.Split(p)
	t, err := about.SearchTree(path.Slug, lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ServeAbout(s, w, r, a, t)
}
