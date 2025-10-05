package auth

import (
	"net/http"

	"g.rg-s.com/sacer/go/requests/meta"
	"g.rg-s.com/sacer/go/viewer"
)

func Route(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	path := m.Path

	vr := "/api/login/verify"
	if len(vr) < len(path) && path[:len(vr)] == vr {
		return VerifyLogin(v, m)
	}

	switch path {
	case "/api/login/request":
		return RequestLogin(v, m)
	case "/api/subscribe":
		return Subscribe(v, m)
	case "/api/register":
		return Register(v, m)
	case "/subscribe", "/login", "/register", "/account":
		return SysPage(v, m)
	default:
		return notFound()
	}
}

func notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}
}
