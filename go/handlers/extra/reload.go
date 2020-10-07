package extra

import (
	"log"
	"net/http"
	"sacer/go/server"
	"sacer/go/server/auth"
)

func AddSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"/", 301)
}

func Reload(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	err := s.LoadSafe()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
}
