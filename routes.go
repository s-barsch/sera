package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"st/pkg/handlers/about"
	"st/pkg/handlers/extra"
	"st/pkg/handlers/front"
	"st/pkg/handlers/graph"
	"st/pkg/handlers/index"
	"st/pkg/handlers/sitemaps"
	"st/pkg/server"
)

func routes(s *server.Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/rl/", makeHandler(s, extra.Reload))

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

	r.HandleFunc("/index", extra.AddSlash)
	r.HandleFunc("/graph", extra.AddSlash)
	r.HandleFunc("/ueber", extra.AddSlash)
	r.HandleFunc("/about", extra.AddSlash)
	r.HandleFunc("/impressum", extra.AddSlash)

	r.PathPrefix("/static/").HandlerFunc(makeHandler(s, extra.StaticFiles))
	r.PathPrefix("/js/").HandlerFunc(makeHandler(s, extra.JSFiles))
	r.HandleFunc("/robots.txt", makeHandler(s, extra.RobotsFiles))

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
