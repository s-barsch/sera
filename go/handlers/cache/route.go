package cache

import (
	"net/http"
	"strconv"

	"g.rg-s.com/sacer/go/handlers/extra"
	"g.rg-s.com/sacer/go/server/meta"
	"g.rg-s.com/sacer/go/viewer"
)

func Route(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	switch rel := m.Path[len("/de/cache"):]; {
	case rel == "/":
		return func() http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/cache", http.StatusMovedPermanently)
			}
		}()
	case rel == "":
		return Main(v, m)

	case isYearPage(m.Split.Slug):
		return Year(v, m)

	case m.Split.IsFile():
		return extra.ServeFile(v, m)

	default:
		return ServeSingle(v, m)
	}
}

func isYearPage(str string) bool {
	if len(str) != 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
