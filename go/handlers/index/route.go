package index

import (
	"net/http"
	p "path/filepath"
	//"stferal/go/entry"
	"stferal/go/handlers/extra"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	reqPath, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := reqPath[len("/index"):]

	if rel == "" {
		Main(s, w, r)
		return
	}

	if rel == "/serial" {
		Serial(s, w, r)
		return
	}

	if rel == "/map.svg" {
		MapSVG(s, w, r)
		return
	}

	lang := head.Lang(r.Host)
	path := paths.Split(reqPath)

	if path.IsFile() {
		extra.ServeFile(s, w, r, path)
		return
	}

	index := s.Trees["index"].Local(s.Flags.Local)[lang]

	if path.Hash == "" {
		t, err := index.SearchTree(path.Slug, lang)
		if err != nil {
			s.Log.Println(err)
			http.NotFound(w, r)
			return
		}
		IndexPage(s, w, r, t)
		return
	}

	t, err := index.LookupTreeHash(path.Hash)
	if err != nil {
		http.Redirect(w, r, p.Dir(r.URL.Path), 301)
		return
	}

	IndexPage(s, w, r, t)
}
