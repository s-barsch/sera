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
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	rel := p[len("/graph"):]

	if rel == "" {
		Main(s, w, r)
		return
	}

	/*
		if rel == "/check" {
			Check(s, w, r)
			return
		}
		*/

		path := paths.Split(p)

		if path.IsFile() {
			extra.ServeFile(s, w, r, path)
			return
		}

		if isYearPage(path.Slug) {
			YearPage(s, w, r, path)
			return
		}

		/*
		El(s, w, r, p)
	*/
}

func isYearPage(str string) bool {
	if len(str) > 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
