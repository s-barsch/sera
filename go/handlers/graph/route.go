package graph

import (
	"net/http"
	"strconv"

	"g.rg-s.com/sera/go/handlers/extra"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
	"g.rg-s.com/sera/go/viewer"
)

func Route(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch {
		case m.Path == "/en/graph" || m.Path == "/de/graph":
			// Main is currently not served directly.
			MainRedirect(w, r, m)

		case m.Split.IsFile():
			extra.ServeFile(w, r, m)

		case isYearPage(m.Split.Slug):
			YearRedirect(w, r, m)

		case isMonth(m.Split.Slug):
			MonthPage(w, r, m)

		default:
			ServeSingle(w, r, m)
		}
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

func Rewrites(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/de"+m.Path, http.StatusMovedPermanently)
	}
}
