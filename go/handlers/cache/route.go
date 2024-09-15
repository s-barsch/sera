package cache

import (
	"net/http"
	"strconv"

	"g.rg-s.com/sferal/go/handlers/extra"
	"g.rg-s.com/sferal/go/server"
	"g.rg-s.com/sferal/go/server/meta"
	"g.rg-s.com/sferal/go/server/paths"
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
