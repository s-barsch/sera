package graph

import (
	"net/http"
	"strconv"

	"g.rg-s.com/sera/go/handlers/extra"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
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

func Route(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	path := paths.Split(p)

	switch {
	case p == "/en/graph" || p == "/de/graph":
		MainRedirect(w, r, m)
		//Main(s, w, r, a)

	case path.IsFile():
		extra.ServeFile(w, r, m, path)

	case isYearPage(path.Slug):
		YearRedirect(w, r, m, path)

	case isMonth(path.Slug):
		MonthPage(w, r, m, path)

	default:
		ServeSingle(w, r, m, path)
	}
}

func isMonth(str string) bool {
	if paths.IsMergedMonths(str) {
		return true
	}
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

func Rewrites(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	http.Redirect(w, r, "/de"+m.Path, http.StatusMovedPermanently)
}
