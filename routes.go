package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"stferal/pkg/handlers/about"
	"stferal/pkg/handlers/extra"
	"stferal/pkg/handlers/front"
	"stferal/pkg/handlers/graph"
	"stferal/pkg/handlers/index"
	"stferal/pkg/handlers/sitemaps"
	"stferal/pkg/server"
)

func routes(s *server.Server) *mux.Router {
	r := mux.NewRouter()

	if s.Flags.Reload {
		r.Use(makeMiddleware(s, extra.ConstantReload))
	}

	r.HandleFunc("/", makeHandler(s, front.Main))

	r.PathPrefix("/index/").HandlerFunc(makeHandler(s, index.Route))
	r.PathPrefix("/graph/").HandlerFunc(makeHandler(s, graph.Route))
	r.PathPrefix("/ueber/").HandlerFunc(makeHandler(s, about.Route))
	r.PathPrefix("/about/").HandlerFunc(makeHandler(s, about.Route))
	r.PathPrefix("/part/").HandlerFunc(makeHandler(s, graph.ElPart))

	r.HandleFunc("/sitemaps.xml", makeHandler(s, sitemaps.Route))
	r.PathPrefix("/sitemaps/").HandlerFunc(makeHandler(s, sitemaps.Route))

	r.PathPrefix("/alt-text/").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/legal/").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/impressum/").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/datenschutz/").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/privacy/").HandlerFunc(makeHandler(s, extra.Route))

	r.PathPrefix("/opt/nachtmodus/").HandlerFunc(makeHandler(s, extra.NightMode))
	r.PathPrefix("/opt/nightmode/").HandlerFunc(makeHandler(s, extra.NightMode))
	r.PathPrefix("/opt/tagmodus/").HandlerFunc(makeHandler(s, extra.DayMode))
	r.PathPrefix("/opt/daymode/").HandlerFunc(makeHandler(s, extra.DayMode))

	r.PathPrefix("/opt/grossschrift/").HandlerFunc(makeHandler(s, extra.LargeType))
	r.PathPrefix("/opt/largetype/").HandlerFunc(makeHandler(s, extra.LargeType))
	r.PathPrefix("/opt/defaulttype/").HandlerFunc(makeHandler(s, extra.DefaultType))
	r.PathPrefix("/opt/standardschrift/").HandlerFunc(makeHandler(s, extra.DefaultType))

	r.PathPrefix("/nolog/").HandlerFunc(makeHandler(s, extra.NoLog))
	r.PathPrefix("/dolog/").HandlerFunc(makeHandler(s, extra.DoLog))

	r.HandleFunc("/index", extra.AddSlash)
	r.HandleFunc("/graph", extra.AddSlash)
	r.HandleFunc("/ueber", extra.AddSlash)
	r.HandleFunc("/about", extra.AddSlash)
	r.HandleFunc("/impressum", extra.AddSlash)

	r.HandleFunc("/rl/", makeHandler(s, extra.Reload))

	r.PathPrefix("/static/").HandlerFunc(makeHandler(s, extra.StaticFiles))
	r.PathPrefix("/js/").HandlerFunc(makeHandler(s, extra.JSFiles))
	r.HandleFunc("/service-worker.js", makeHandler(s, extra.ServiceWorker))
	r.HandleFunc("/robots.txt", makeHandler(s, extra.RobotsFiles))

	r.PathPrefix("/manifest.json").HandlerFunc(makeHandler(s, extra.Manifest))
	r.PathPrefix("/manifest-night.json").HandlerFunc(makeHandler(s, extra.Manifest))

	fileRoutes := map[string]string{
		"/googledbd0f1dfe416dbee.html": "/static/seo/googledbd0f1dfe416dbee.html",
		"/BingSiteAuth.xml":            "/static/seo/BingSiteAuth.xml",
		"/manifest-de.webmanifest":     "/static/manifest-de.webmanifest",
		"/manifest-en.webmanifest":     "/static/manifest-en.webmanifest",
	}

	for query := range fileRoutes {
		path := fileRoutes[query]
		r.HandleFunc(query, func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = path
			extra.StaticFiles(s, w, r)
		})
	}

	return r
}

func makeHandler(s *server.Server, fn func(*server.Server, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(s, w, r)
	}
}

func makeMiddleware(s *server.Server, fn func(*server.Server, http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return fn(s, next)
	}
}
