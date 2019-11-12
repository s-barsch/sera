package extra

import (
	"net/http"
	"stferal/pkg/server"
)

func LargeType(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "largetype",
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

func DefaultType(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "largetype",
		Path:   "/",
		MaxAge: -1,
	})
	ref := r.Referer()
	if ref == "" {
		ref = "/"
	}
	http.Redirect(w, r, ref, 307)
}
