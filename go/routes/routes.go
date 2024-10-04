package routes

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/handlers/about"
	"g.rg-s.com/sera/go/handlers/auth"
	"g.rg-s.com/sera/go/handlers/cache"
	"g.rg-s.com/sera/go/handlers/extra"
	"g.rg-s.com/sera/go/handlers/front"
	"g.rg-s.com/sera/go/handlers/graph"
	"g.rg-s.com/sera/go/server"

	//"g.rg-s.com/sera/go/handlers/indecs"
	//"g.rg-s.com/sera/go/handlers/index"

	//"g.rg-s.com/sera/go/handlers/sitemaps"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"

	"github.com/gorilla/mux"
)

func Router(s *server.Server) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", makeHandler(front.Main))
	r.HandleFunc("/en", makeHandler(front.Rewrites))
	r.HandleFunc("/en/", makeHandler(front.Rewrites))
	r.HandleFunc("/de", makeHandler(front.Main))
	r.PathPrefix("/de/graph").HandlerFunc(makeHandler(graph.Route))
	r.PathPrefix("/en/graph").HandlerFunc(makeHandler(graph.Route))
	r.PathPrefix("/de/cache").HandlerFunc(makeHandler(cache.Route))
	r.PathPrefix("/en/cache").HandlerFunc(makeHandler(cache.Route))
	r.PathPrefix("/de/ueber").HandlerFunc(makeHandler(about.About))
	r.PathPrefix("/de/about").HandlerFunc(makeHandler(about.About))
	r.PathPrefix("/en/about").HandlerFunc(makeHandler(about.About))

	r.PathPrefix("/ueber").HandlerFunc(makeHandler(about.Rewrites))
	r.PathPrefix("/about").HandlerFunc(makeHandler(about.Rewrites))
	r.PathPrefix("/graph").HandlerFunc(makeHandler(graph.Rewrites))

	/*
		r.PathPrefix("/indecs").HandlerFunc(makeHandler(s, indecs.Route))
		r.PathPrefix("/index").HandlerFunc(makeHandler(s, index.Route))
		r.PathPrefix("/part/").HandlerFunc(makeHandler(s, graph.ElPart))
	*/

	r.PathPrefix("/api").HandlerFunc(makeHandler(auth.Route))
	r.PathPrefix("/register").HandlerFunc(makeHandler(auth.Route))
	r.PathPrefix("/subscribe").HandlerFunc(makeHandler(auth.Route))
	r.PathPrefix("/login").HandlerFunc(makeHandler(auth.Route))
	r.PathPrefix("/account").HandlerFunc(makeHandler(auth.Route))

	/*
		r.HandleFunc("/sitemaps.xml", makeHandler(s, sitemaps.Route))
		r.PathPrefix("/sitemaps").HandlerFunc(makeHandler(s, sitemaps.Route))
	*/

	r.PathPrefix("/de/impressum").HandlerFunc(makeHandler(extra.Extra))
	r.PathPrefix("/legal").HandlerFunc(makeHandler(extra.Extra))
	r.PathPrefix("/de/datenschutz").HandlerFunc(makeHandler(extra.Extra))
	r.PathPrefix("/privacy").HandlerFunc(makeHandler(extra.Extra))

	r.HandleFunc("/opt/{option}/{value}", makeHandler(extra.SetOption))

	r.HandleFunc("/rl/", makeHandler(extra.Reload))

	r.PathPrefix("/static/").HandlerFunc(makeHandler(extra.StaticFiles))
	r.PathPrefix("/js/").HandlerFunc(makeHandler(extra.JSFiles))
	r.HandleFunc("/sw.js", makeHandler(extra.ServiceWorker))
	r.HandleFunc("/robots.txt", makeHandler(extra.RobotsFiles))

	r.PathPrefix("/de/manifest.json").HandlerFunc(makeHandler(extra.Manifest))
	r.PathPrefix("/manifest.json").HandlerFunc(makeHandler(extra.Manifest))

	fileRoutes := map[string]string{
		"/BingSiteAuth.xml": "/static/seo/BingSiteAuth.xml",
	}

	for query := range fileRoutes {
		path := fileRoutes[query]
		r.HandleFunc(query, func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = path
			m, err := meta.NewMeta(s.Users, w, r)
			if err != nil {
				log.Println(err)
			}

			extra.StaticFiles(w, r, m)
		})
	}

	return r
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, *meta.Meta)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := meta.NewMeta(s.Store.Users, w, r)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		fn(w, r, m)
	}
}
