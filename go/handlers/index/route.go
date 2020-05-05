package index

import (
	"net/http"
	p "path/filepath"
	//"stferal/go/entry"
	//"stferal/go/handlers/extra"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := path[len("/index"):]

	if rel == "" {
		Main(s, w, r)
		return
	}

	/*
	if rel == "/serial/" {
		Serial(s, w, r)
		return
	}

	if rel == "/map.svg" {
		MapSVG(s, w, r)
		return
	}
	*/

	pa := paths.Split(path)

	/*
	if p.Subdir != "" {
		extra.Files(s, w, r, p)
		return
	}
	*/

	index := s.Trees["index"]

	/*
	if s.Flags.Local {
		tree = s.Trees["index-private"]
	}
	*/

	if pa.Hash == "" {
		t, err := index.SearchTree(pa.Slug, head.Lang(r.Host))
		if err != nil {
			s.Log.Println(err)
			http.NotFound(w, r)
			return
		}
		IndexPage(s, w, r, t)
		return
	}

	t, err := index.LookupTreeHash(pa.Hash)
	if err != nil {
		http.Redirect(w, r, p.Dir(r.URL.Path), 301)
		return
	}

	IndexPage(s, w, r, t)
}
