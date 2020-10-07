package kine 

import (
	"sacer/go/paths"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/handlers/extra"
	"net/http"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
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
		Main(s, w, r, a)
		return
	}

	path := paths.Split(p)

	if path.IsFile() {
		extra.ServeFile(s, w, r, path)
		return
	}

	ServeSingle(s, w, r, path)
}
