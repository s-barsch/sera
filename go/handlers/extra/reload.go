package extra

import (
	"log"
	"net/http"

	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

func AddSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
}

func Reload(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.Srv.LoadSafe()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
