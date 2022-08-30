package routes

import (
	"log"
	"net/http"
	"sacer/go/handlers/about"
	authH "sacer/go/handlers/auth"
	"sacer/go/handlers/extra"
	"sacer/go/handlers/front"
	"sacer/go/handlers/graph"
	"sacer/go/handlers/indecs"
	"sacer/go/handlers/index"
	"sacer/go/handlers/kine"
	"sacer/go/handlers/sitemaps"
	"sacer/go/server"
	"sacer/go/server/users"

	"github.com/gorilla/mux"
)

func Router(s *server.Server) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", makeHandler(s, front.Main))
	r.PathPrefix("/index").HandlerFunc(makeHandler(s, index.Route))
	r.PathPrefix("/graph").HandlerFunc(makeHandler(s, graph.Route))
	r.PathPrefix("/kine").HandlerFunc(makeHandler(s, kine.Route))
	r.PathPrefix("/cine").HandlerFunc(makeHandler(s, kine.Route))
	r.PathPrefix("/ueber").HandlerFunc(makeHandler(s, about.Route))
	r.PathPrefix("/about").HandlerFunc(makeHandler(s, about.Route))

	r.PathPrefix("/indecs").HandlerFunc(makeHandler(s, indecs.Route))
	/*
		r.PathPrefix("/part/").HandlerFunc(makeHandler(s, graph.ElPart))
	*/

	r.PathPrefix("/api").HandlerFunc(makeHandler(s, authH.Route))
	r.PathPrefix("/register").HandlerFunc(makeHandler(s, authH.Route))
	r.PathPrefix("/subscribe").HandlerFunc(makeHandler(s, authH.Route))
	r.PathPrefix("/login").HandlerFunc(makeHandler(s, authH.Route))

	r.HandleFunc("/sitemaps.xml", makeHandler(s, sitemaps.Route))
	r.PathPrefix("/sitemaps").HandlerFunc(makeHandler(s, sitemaps.Route))


	r.PathPrefix("/legal").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/impressum").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/datenschutz").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/privacy").HandlerFunc(makeHandler(s, extra.Route))

	r.HandleFunc("/opt/{option}/{value}", makeHandler(s, extra.SetOption))

	r.HandleFunc("/rl/", makeHandler(s, extra.Reload))

	r.PathPrefix("/static/").HandlerFunc(makeHandler(s, extra.StaticFiles))
	r.PathPrefix("/js/").HandlerFunc(makeHandler(s, extra.JSFiles))
	r.HandleFunc("/sw.js", makeHandler(s, extra.ServiceWorker))
	r.HandleFunc("/robots.txt", makeHandler(s, extra.RobotsFiles))

	r.PathPrefix("/manifest.json").HandlerFunc(makeHandler(s, extra.Manifest))

	fileRoutes := map[string]string{
		"/BingSiteAuth.xml": "/static/seo/BingSiteAuth.xml",
	}

	for query := range fileRoutes {
		path := fileRoutes[query]
		r.HandleFunc(query, func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = path
			a, err := s.Users.CheckAuth(r)
			if err != nil && err != http.ErrNoCookie {
				log.Println(err)
			}
			extra.StaticFiles(s, w, r, a)
		})
	}

	return r
}

func makeHandler(s *server.Server, fn func(*server.Server, http.ResponseWriter, *http.Request, *users.Auth)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := s.Users.CheckAuth(r)
		if err != nil {
			log.Println(err)
		}
		fn(s, w, r, a)
	}
}
