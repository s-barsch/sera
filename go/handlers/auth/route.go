package auth

import (
	"net/http"
	"sacer/go/server/paths"
	"sacer/go/server"
	"sacer/go/server/auth"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch path {
	case "/api/login":
		Login(w, r)
	case "/api/subscribe":
		Subscribe(w, r)
	case "/api/register":
		Register(w, r)
	case "/subscribe", "/login", "/register":
		SysPage(s, w, r, a)
	default:
		http.NotFound(w, r)
	}
}
