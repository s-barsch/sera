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

	lang := head.Lang(r.Host)

	rel := p[len("/about"):] // same length as "ueber"

	if rel == "" {
		ServeAbout(s, w, r, s.Trees["about"].Public[lang])
		return
	}

	path := paths.Split(p)
	about := s.Trees["about"].Public[lang]

	t, err := about.SearchTree(path.Slug, lang)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ServeAbout(s, w, r, t)
}

