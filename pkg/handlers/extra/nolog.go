package extra

import (
	"net/http"
	"stferal/pkg/server"
)

func NoLog(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "nolog",
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

func DoLog(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "nolog",
		Path:   "/",
		MaxAge: -1,
	})
	ref := r.Referer()
	if ref == "" {
		ref = "/"
	}
	http.Redirect(w, r, ref, 307)
}
