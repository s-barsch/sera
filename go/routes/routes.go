package routes

import (
	"net/http"
	"sacer/go/handlers/about"
	"sacer/go/handlers/auth"
	"sacer/go/handlers/extra"
	"sacer/go/handlers/front"
	"sacer/go/handlers/graph"

	//"sacer/go/handlers/indecs"
	//"sacer/go/handlers/index"
	"sacer/go/handlers/kine"
	//"sacer/go/handlers/sitemaps"
	"sacer/go/server"
	"sacer/go/server/meta"

	"github.com/gorilla/mux"
)

func Router(s *server.Server) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", makeHandler(s, front.Main))
	r.HandleFunc("/en", makeHandler(s, front.Rewrites))
	r.HandleFunc("/en/", makeHandler(s, front.Rewrites))
	r.HandleFunc("/de", makeHandler(s, front.Main))
	r.PathPrefix("/de/graph").HandlerFunc(makeHandler(s, graph.Route))
	r.PathPrefix("/en/graph").HandlerFunc(makeHandler(s, graph.Route))
	r.PathPrefix("/de/kine").HandlerFunc(makeHandler(s, kine.Route))
	r.PathPrefix("/en/cine").HandlerFunc(makeHandler(s, kine.Route))
	r.PathPrefix("/de/ueber").HandlerFunc(makeHandler(s, about.Route))
	r.PathPrefix("/en/about").HandlerFunc(makeHandler(s, about.Route))

	r.PathPrefix("/ueber").HandlerFunc(makeHandler(s, about.Rewrites))
	r.PathPrefix("/about").HandlerFunc(makeHandler(s, about.Rewrites))
	r.PathPrefix("/kine").HandlerFunc(makeHandler(s, kine.Rewrites))
	r.PathPrefix("/cine").HandlerFunc(makeHandler(s, kine.Rewrites))
	r.PathPrefix("/graph").HandlerFunc(makeHandler(s, graph.Rewrites))

	/*
		r.PathPrefix("/indecs").HandlerFunc(makeHandler(s, indecs.Route))
		r.PathPrefix("/index").HandlerFunc(makeHandler(s, index.Route))
		r.PathPrefix("/part/").HandlerFunc(makeHandler(s, graph.ElPart))
	*/

	r.PathPrefix("/api").HandlerFunc(makeHandler(s, auth.Route))
	r.PathPrefix("/register").HandlerFunc(makeHandler(s, auth.Route))
	r.PathPrefix("/subscribe").HandlerFunc(makeHandler(s, auth.Route))
	r.PathPrefix("/login").HandlerFunc(makeHandler(s, auth.Route))
	r.PathPrefix("/account").HandlerFunc(makeHandler(s, auth.Route))

	/*
		r.HandleFunc("/sitemaps.xml", makeHandler(s, sitemaps.Route))
		r.PathPrefix("/sitemaps").HandlerFunc(makeHandler(s, sitemaps.Route))
	*/

	r.PathPrefix("/de/impressum").HandlerFunc(makeHandler(s, extra.Extra))
	r.PathPrefix("/legal").HandlerFunc(makeHandler(s, extra.Extra))
	r.PathPrefix("/de/datenschutz").HandlerFunc(makeHandler(s, extra.Extra))
	r.PathPrefix("/privacy").HandlerFunc(makeHandler(s, extra.Extra))

	r.HandleFunc("/opt/{option}/{value}", makeHandler(s, extra.SetOption))

	r.HandleFunc("/rl/", makeHandler(s, extra.Reload))

	r.PathPrefix("/static/").HandlerFunc(makeHandler(s, extra.StaticFiles))
	r.PathPrefix("/js/").HandlerFunc(makeHandler(s, extra.JSFiles))
	r.HandleFunc("/sw.js", makeHandler(s, extra.ServiceWorker))
	r.HandleFunc("/robots.txt", makeHandler(s, extra.RobotsFiles))

	r.PathPrefix("/de/manifest.json").HandlerFunc(makeHandler(s, extra.Manifest))
	r.PathPrefix("/manifest.json").HandlerFunc(makeHandler(s, extra.Manifest))

	fileRoutes := map[string]string{
		"/BingSiteAuth.xml": "/static/seo/BingSiteAuth.xml",
	}

	for query := range fileRoutes {
		path := fileRoutes[query]
		r.HandleFunc(query, func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = path
			m, err := meta.NewMeta(s.Users, w, r)
			if err != nil {
				s.Log.Println(err)
			}

			extra.StaticFiles(s, w, r, m)
		})
	}

	return r
}

func makeHandler(s *server.Server, fn func(*server.Server, http.ResponseWriter, *http.Request, *meta.Meta)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := meta.NewMeta(s.Users, w, r)
		if err != nil {
			s.Log.Println(err)
			http.Error(w, "internal error", 502)
			return
		}
		fn(s, w, r, m)
	}
}
