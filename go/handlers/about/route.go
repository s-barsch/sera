package about

import (
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func Rewrites(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	folder := m.Path[:len("/about")]
	if folder == "/about" {
		http.Redirect(w, r, "/en"+m.Path, http.StatusMovedPermanently)
		return
	}
	/*
		if folder == "/ueber" {
			http.Redirect(w, r, "/de"+m.Path, http.StatusMovedPermanently)
			return
		}
	*/
}
