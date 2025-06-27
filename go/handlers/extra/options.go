package extra

import (
	"net/http"

	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
	"github.com/gorilla/mux"
)

const expire = 60 * 60 * 24 * 365 // 1 year

func SetOption(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		http.SetCookie(w, &http.Cookie{
			Name:   v["option"],
			Value:  v["value"],
			Path:   "/",
			MaxAge: expire,
		})
		http.Redirect(w, r, toRef(r.Referer()), http.StatusTemporaryRedirect)
	}
}

func toRef(path string) string {
	if path == "" {
		return "/"
	}
	return path
}
