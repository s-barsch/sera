package kine

import (
	"net/http"
	"sacer/go/handlers/extra"
	"sacer/go/server"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := p[len("/de/kine"):]

	if rel == "/" {
		http.Redirect(w, r, "/kine", 301)
		return
	}

	if rel == "" {
		Main(s, w, r, m)
		return
	}

	path := paths.Split(p)

	if path.IsFile() {
		extra.ServeFile(s, w, r, m, path)
		return
	}

	ServeSingle(s, w, r, m, path)
}

func Rewrites(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	folder := m.Path[:len("/kine")]
	if folder == "/cine" {
		http.Redirect(w, r, "/en"+m.Path, 301)
		return
	}
	if folder == "/kine" {
		http.Redirect(w, r, "/de"+m.Path, 301)
		return
	}
}
