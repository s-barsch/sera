package video

import (
	"stferal/go/paths"
	"stferal/go/server"
	"stferal/go/handlers/extra"
	"net/http"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	p, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	path := paths.Split(p)

	if path.IsFile() {
		extra.ServeFile(s, w, r, path)
		return
	}

}
