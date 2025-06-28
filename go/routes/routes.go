package routes

import (
	"net/http"

	"g.rg-s.com/sera/go/handlers/about"
	"g.rg-s.com/sera/go/handlers/auth"
	"g.rg-s.com/sera/go/handlers/cache"
	"g.rg-s.com/sera/go/handlers/extra"
	"g.rg-s.com/sera/go/handlers/front"
	"g.rg-s.com/sera/go/handlers/graph"
	"g.rg-s.com/sera/go/viewer"

	//"g.rg-s.com/sera/go/handlers/indecs"
	//"g.rg-s.com/sera/go/handlers/index"

	//"g.rg-s.com/sera/go/handlers/sitemaps"

	"github.com/gorilla/mux"
)

type router struct {
	viewer *viewer.Viewer
	mux    *mux.Router
}

func newRouter(v *viewer.Viewer) *router {
	return &router{
		viewer: v,
		mux:    mux.NewRouter().StrictSlash(true),
	}
}

func (r *router) Mux() http.Handler {
	return r.mux
}

func (r *router) register(path string, f viewer.HandleFunc) {
	r.mux.PathPrefix(path).HandlerFunc(r.viewer.View(f))
}

func (r *router) registerExact(path string, f viewer.HandleFunc) {
	r.mux.HandleFunc(path, r.viewer.View(f))
}

func Router(v *viewer.Viewer) http.Handler {
	r := newRouter(v)

	r.registerExact("/", front.Main)
	r.registerExact("/en", front.Rewrites)
	r.registerExact("/en/", front.Rewrites)
	r.registerExact("/de", front.Main)
	r.register("/de/graph", graph.Route)
	r.register("/en/graph", graph.Route)
	r.register("/de/cache", cache.Route)
	r.register("/en/cache", cache.Route)
	r.register("/de/ueber", about.About)
	r.register("/de/about", about.About)
	r.register("/en/about", about.About)

	r.register("/ueber", about.Rewrites)
	r.register("/about", about.Rewrites)
	r.register("/graph", graph.Rewrites)

	/*
		r.register("/indecs", s, indecs.Route))
		r.register("/index", s, index.Route))
		r.register("/part/", s, graph.ElPart))
	*/

	r.register("/api", auth.Route)
	r.register("/register", auth.Route)
	r.register("/subscribe", auth.Route)
	r.register("/login", auth.Route)
	r.register("/account", auth.Route)

	/*
		r.registerExact("/sitemaps.xml", makeHandler(s, sitemaps.Route))
		r.register("/sitemaps", s, sitemaps.Route))
	*/

	r.register("/de/impressum", extra.Extra)
	r.register("/legal", extra.Extra)
	r.register("/de/datenschutz", extra.Extra)
	r.register("/privacy", extra.Extra)

	r.registerExact("/opt/{option}/{value}", extra.SetOption)

	r.registerExact("/rl/", extra.Reload)

	r.register("/static/", extra.StaticFiles)
	r.register("/js/", extra.JSFiles)
	r.registerExact("/sw.js", extra.ServiceWorker)
	r.registerExact("/robots.txt", extra.RobotsFiles)

	r.register("/de/manifest.json", extra.Manifest)
	r.register("/manifest.json", extra.Manifest)

	return r.Mux()
}
