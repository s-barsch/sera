package extra

import (
	"net/http"
	"sacer/go/server"
)

const expire = 60 * 60 * 24 * 365 // 1 year

func DarkColors(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "colors",
		Value:  "dark",
		Path:   "/",
		MaxAge: expire,
	})
	http.Redirect(w, r, toRef(r.Referer()), 307)
}

func LightColors(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "colors",
		Value:  "light",
		Path:   "/",
		MaxAge: expire,
	})
	http.Redirect(w, r, toRef(r.Referer()), 307)
}

func LargeType(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "type",
		Value:  "large",
		Path:   "/",
		MaxAge: expire,
	})
	http.Redirect(w, r, toRef(r.Referer()), 307)
}

func SmallType(s *server.Server, w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "type",
		Value:  "small",
		Path:   "/",
		MaxAge: expire,
	})
	http.Redirect(w, r, toRef(r.Referer()), 307)
}

func toRef(path string) string {
	if path == "" {
		return "/"
	}
	return path
}
