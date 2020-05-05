package graph

import (
	"net/http"
	"stferal/go/handlers/extra"
	"stferal/go/paths"
	"stferal/go/server"
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

	rel := r.URL.Path[len("/graph"):]

	if rel == "" {
		Main(s, w, r)
		return
	}

	/*
	if rel == "/check" {
		Check(s, w, r)
		return
	}

	p := paths.Split(path)

	if p.Subdir != "" {
		extra.Files(s, w, r, p)
		return
	}

	if isTimePage(p.Acronym) {
		Year(s, w, r, p)
		return
	}

	El(s, w, r, p)
	*/
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
