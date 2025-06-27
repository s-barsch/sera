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
	switch {
	case m.Path == "/en/graph" || m.Path == "/de/graph":
		// Main is currently not served directly.
		return MainRedirect(v, m)

	case m.Split.IsFile():
		return extra.ServeFile(v, m)

	case isYearPage(m.Split.Slug):
		return YearRedirect(v, m)

	case isMonth(m.Split.Slug):
		return MonthPage(v, m)

	default:
		return ServeSingle(v, m)
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
