package cache

import (
	"net/http"
	"strconv"

	"g.rg-s.com/sera/go/handlers/extra"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

func Route(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch rel := m.Path[len("/de/cache"):]; {

		case rel == "/":
			http.Redirect(w, r, "/cache", http.StatusMovedPermanently)

		case rel == "":
			Main(w, r, m)

		case isYearPage(m.Split.Slug):
			Year(w, r, m)

		case m.Split.IsFile():
			extra.ServeFile(w, r, m)

		default:
			ServeSingle(w, r, m)
		}
	}
}

func isYearPage(str string) bool {
	if len(str) != 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
