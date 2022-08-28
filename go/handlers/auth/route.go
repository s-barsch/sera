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

	vr := "/api/login/verify"
	if len(vr) < len(path) && path[:len(vr)] == vr {
		VerifyLogin(w, r)
		return
	}

	switch path {
	case "/api/login/request":
		RequestLogin(w, r)
	case "/api/subscribe":
		Subscribe(w, r)
	case "/api/register":
		Register(w, r)
	case "/subscribe", "/login", "/register":
		SysPage(s, w, r, a)
	default:
		println(path)
		http.NotFound(w, r)
	}
}

