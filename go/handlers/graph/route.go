package graph

import (
	"net/http"
	"sacer/go/handlers/extra"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/paths"
	"strconv"
)

/*
func graphPart(w http.ResponseWriter, r *http.Request) {
	serveGraphElementPart(w, r, splitPath(r.URL.Path))
}
*/

/*
	if rel == "/check" {
		Check(s, w, r)
		return
	}
*/

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	path := paths.Split(p)

	switch {

	case p == "/graph":
		Main(s, w, r, a)

	case path.IsFile():
		extra.ServeFile(s, w, r, a, path)

	case isYearPage(path.Slug):
		if !s.Flags.Local {
			http.Error(w, "temporarily unavailable", 503)
			return
		}
		YearPage(s, w, r, a, path)

	case isMonth(path.Slug):
		MonthPage(s, w, r, a, path)

	default:
		ServeSingle(s, w, r, a, path)
	}
}

func isMonth(str string) bool {
	if len(str) != 2 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}

func isYearPage(str string) bool {
	if len(str) != 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
