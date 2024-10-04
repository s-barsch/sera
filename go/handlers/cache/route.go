package cache

import (
	"net/http"
	"strconv"

	"g.rg-s.com/sera/go/handlers/extra"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
)

func Route(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	rel := p[len("/de/cache"):]

	if rel == "/" {
		http.Redirect(w, r, "/cache", http.StatusMovedPermanently)
		return
	}

	if rel == "" {
		Main(w, r, m)
		return
	}
	path := paths.Split(p)

	if isYearPage(path.Slug) {
		Year(w, r, m, path)
		return
	}

	if path.IsFile() {
		extra.ServeFile(w, r, m, path)
		return
	}

	ServeSingle(w, r, m, path)
}

func isYearPage(str string) bool {
	if len(str) != 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
