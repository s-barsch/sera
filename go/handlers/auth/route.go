package auth

import (
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
	"g.sacerb.com/sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	vr := "/api/login/verify"
	if len(vr) < len(path) && path[:len(vr)] == vr {
		VerifyLogin(s, w, r)
		return
	}

	switch path {
	case "/api/login/request":
		RequestLogin(s, w, r)
	case "/api/subscribe":
		Subscribe(s, w, r)
	case "/api/register":
		Register(s, w, r)
	case "/subscribe", "/login", "/register", "/account":
		SysPage(s, w, r, m)
	default:
		http.NotFound(w, r)
	}
}
