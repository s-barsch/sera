package extra

import (
	"log"
	"net/http"
	"sacer/go/server"
	//"sacer/go/handlers/index"
	p "path/filepath"
)

func AddSlash(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"/", 301)
}

func ConstantReload(s *server.Server, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p.Ext(r.URL.Path) == "" {
			log.Println(r.URL.Path)
			err := s.Load()
			if err != nil {
				log.Println(err)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func Reload(s *server.Server, w http.ResponseWriter, r *http.Request) {
	err := s.Load()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	}

	//index.SaveMaps(s)
}
