package extra

import (
	"net/http"

	"g.rg-s.com/sferal/go/server"
	"g.rg-s.com/sferal/go/server/meta"
	"github.com/gorilla/mux"
)

const expire = 60 * 60 * 24 * 365 // 1 year

func SetOption(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	v := mux.Vars(r)
	http.SetCookie(w, &http.Cookie{
		Name:   v["option"],
		Value:  v["value"],
		Path:   "/",
		MaxAge: expire,
	})
	http.Redirect(w, r, toRef(r.Referer()), http.StatusTemporaryRedirect)
}

func toRef(path string) string {
	if path == "" {
		return "/"
	}
	return path
}
