package routes 

import (
	"github.com/gorilla/mux"
	"net/http"
	"sacer/go/handlers/index"
	"sacer/go/handlers/about"
	"sacer/go/handlers/extra"
	"sacer/go/handlers/graph"
	"sacer/go/handlers/front"
	"sacer/go/handlers/kine"
	"sacer/go/handlers/sitemaps"
	"sacer/go/server"
	"sacer/go/server/auth"
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

	/*
	r.PathPrefix("/part/").HandlerFunc(makeHandler(s, graph.ElPart))
	r.PathPrefix("/alt-text/").HandlerFunc(makeHandler(s, extra.Route))
	*/

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
			extra.StaticFiles(s, w, r, auth.CheckAuth(r))
		})
	}

	return r
}

func makeHandler(s *server.Server, fn func(*server.Server, http.ResponseWriter, *http.Request, *auth.Auth)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a := auth.CheckAuth(r)
		if s.Flags.Local {
			a.Subscriber = true
		}
		fn(s, w, r, a)
	}
}

