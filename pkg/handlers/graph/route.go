package graph

import (
	"net/http"
	"stferal/pkg/handlers/extra"
	"stferal/pkg/paths"
	"stferal/pkg/server"
	"strconv"
)

/*
func graphPart(w http.ResponseWriter, r *http.Request) {
	serveGraphElementPart(w, r, splitPath(r.URL.Path))
}
*/

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if path == "/graph/" {
		Main(s, w, r)
		return
	}

	p := paths.Split(r.URL.Path)

	if isTimePage(p.Acronym) {
		Year(s, w, r, p)
		return
	}

	if p.Type != "" {
		extra.Files(s, w, r, p)
		return
	}

	if path == "/graph/check/" {
		Check(s, w, r)
		return
	}

	El(s, w, r, p)
}

func isTimePage(acr string) bool {
	if len(acr) > 4 {
		return false
	}
	_, err := strconv.Atoi(acr)
	if err != nil {
		return false
	}
	return true
}
