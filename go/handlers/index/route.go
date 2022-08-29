package index

import (
	"net/http"
	"sacer/go/server"
	"sacer/go/server/users"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	Main(s, w, r, a)
}
