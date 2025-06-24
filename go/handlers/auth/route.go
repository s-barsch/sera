package auth

import (
	"net/http"

	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
)

func Route(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
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
	case "/subscribe", "/login", "/register", "/account":
		SysPage(w, r, m)
	default:
		http.NotFound(w, r)
	}
}
