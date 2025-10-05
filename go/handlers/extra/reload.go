package extra

import (
	"log"
	"net/http"

	"g.rg-s.com/sacer/go/server/meta"
	"g.rg-s.com/sacer/go/viewer"
)

func AddSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
}

func Reload(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := v.Reload()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
