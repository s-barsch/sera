package komposita

import (
	"net/http"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/paths"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	p := paths.Split(path)


	Article(s, w, r, a, p)
}
