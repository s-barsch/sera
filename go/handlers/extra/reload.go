package extra

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func AddSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"/", http.StatusMovedPermanently)
}

func Reload(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	err := server.Store.LoadSafe()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}
}
