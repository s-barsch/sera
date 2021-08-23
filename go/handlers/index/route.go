package index

import (
	"net/http"
	"sacer/go/server"
	"sacer/go/server/auth"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	Main(s, w, r, a)
}
