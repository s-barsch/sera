package routes 

import (
	"github.com/gorilla/mux"
	"net/http"
	"stferal/go/handlers/about"
	"stferal/go/handlers/extra"
	"stferal/go/handlers/front"
	"stferal/go/handlers/graph"
	"stferal/go/handlers/index"
	"stferal/go/handlers/sitemaps"
	"stferal/go/server"
)

func Router(s *server.Server) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	if s.Flags.Reload {
		r.Use(makeMiddleware(s, extra.ConstantReload))
	}

	r.HandleFunc("/", makeHandler(s, front.Main))

	r.PathPrefix("/index").HandlerFunc(makeHandler(s, index.Route))
	r.PathPrefix("/graph").HandlerFunc(makeHandler(s, graph.Route))
	r.PathPrefix("/ueber").HandlerFunc(makeHandler(s, about.Route))
	r.PathPrefix("/about").HandlerFunc(makeHandler(s, about.Route))

	r.PathPrefix("/part/").HandlerFunc(makeHandler(s, graph.ElPart))
	r.PathPrefix("/alt-text/").HandlerFunc(makeHandler(s, extra.Route))

	r.HandleFunc("/sitemaps.xml", makeHandler(s, sitemaps.Route))
	r.PathPrefix("/sitemaps").HandlerFunc(makeHandler(s, sitemaps.Route))

	r.PathPrefix("/legal").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/impressum").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/datenschutz").HandlerFunc(makeHandler(s, extra.Route))
	r.PathPrefix("/privacy").HandlerFunc(makeHandler(s, extra.Route))

	r.HandleFunc("/opt/colors/dark", makeHandler(s, extra.DarkColors))
	r.HandleFunc("/opt/colors/light", makeHandler(s, extra.LightColors))
	r.HandleFunc("/opt/type/large", makeHandler(s, extra.LargeType))
	r.HandleFunc("/opt/type/small", makeHandler(s, extra.SmallType))

	r.HandleFunc("/rl/", makeHandler(s, extra.Reload))

	r.PathPrefix("/static/").HandlerFunc(makeHandler(s, extra.StaticFiles))
	r.PathPrefix("/js/").HandlerFunc(makeHandler(s, extra.JSFiles))
	r.HandleFunc("/sw.js", makeHandler(s, extra.ServiceWorker))
	r.HandleFunc("/robots.txt", makeHandler(s, extra.RobotsFiles))

	r.PathPrefix("/manifest.json").HandlerFunc(makeHandler(s, extra.Manifest))

	fileRoutes := map[string]string{
		"/googledbd0f1dfe416dbee.html": "/static/seo/googledbd0f1dfe416dbee.html",
		"/BingSiteAuth.xml":            "/static/seo/BingSiteAuth.xml",
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
