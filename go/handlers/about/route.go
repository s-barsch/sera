package about

import (
	"net/http"
	//"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := p[len("/about"):] // same length as "ueber"
	lang := head.Lang(r.Host)
	about := s.Trees["about"].Local(s.Flags.Local)[lang]

	if rel == "" {
		ServeAbout(s, w, r, about)
		return
	}

	path := paths.Split(p)
	t, err := about.SearchTree(path.Slug, lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ServeAbout(s, w, r, t)
}

