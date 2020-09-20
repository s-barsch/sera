package kine 

import (
	"stferal/go/paths"
	"stferal/go/server"
	"stferal/go/handlers/extra"
	"net/http"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := p[len("/kine"):]

	if rel == "/" {
		http.Redirect(w, r, "/kine", 301)
		return
	}

	if rel == "" {
		Main(s, w, r)
		return
	}

	path := paths.Split(p)

	if path.IsFile() {
		extra.ServeFile(s, w, r, path)
		return
	}

	ServeSingle(s, w, r, path)
}
