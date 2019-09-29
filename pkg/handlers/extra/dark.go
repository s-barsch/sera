package extra

import (
	"net/http"
	"stferal/pkg/server"
)

func EnableDarkMode(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "darkmode",
		Value:  "true",
		Path:   "/",
		MaxAge: 60 * 60 * 24 * 365, 
	})
	ref := r.Referer()
	if ref == "" {
		ref = "/"
	}
	http.Redirect(w, r, ref, 307)
}

func DisableDarkMode(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "darkmode",
		Path:   "/",
		MaxAge: -1,
	})
	ref := r.Referer()
	if ref == "" {
		ref = "/"
	}
	http.Redirect(w, r, ref, 307)
}
