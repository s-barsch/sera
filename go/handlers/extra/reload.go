package extra

import (
	"log"
	"net/http"
	"sacer/go/server"
	"sacer/go/server/meta"
)

func AddSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"/", 301)
}

func Reload(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	err := s.LoadSafe()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
}
