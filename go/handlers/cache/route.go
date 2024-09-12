package cache

import (
	"net/http"
	"strconv"

	"g.sacerb.com/sacer/go/handlers/extra"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
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
		Main(s, w, r, m)
		return
	}
	path := paths.Split(p)

	if isYearPage(path.Slug) {
		Year(s, w, r, m, path)
		return
	}

	if path.IsFile() {
		extra.ServeFile(s, w, r, m, path)
		return
	}

	ServeSingle(s, w, r, m, path)
}

func isYearPage(str string) bool {
	if len(str) != 4 {
		return false
	}
	_, err := strconv.Atoi(str)
	return err == nil
}
