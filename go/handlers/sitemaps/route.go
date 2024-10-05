package sitemaps

import (
	"net/http"

	"g.rg-s.com/sera/go/server/meta"
)

func Route(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	switch r.URL.Path {
	case "/sitemaps.xml":
		Index(w, r, m)
	case "/sitemaps/core.xml":
		Core(w, r, m)
	case "/sitemaps/trees.xml":
		Trees(w, r, m)
	case "/sitemaps/caches.xml":
		Kines(w, r, m)
	case "/sitemaps/graph-entries.xml":
		GraphEntries(w, r, m)
	default:
		http.NotFound(w, r)
	}
}

/*
	r.HandleFunc("/sitemaps.xml", sitemapIndex)
	r.HandleFunc("/sitemaps/core.xml", sitemapCore)
	r.HandleFunc("/sitemaps/holds.xml", sitemapHolds)
	r.HandleFunc("/sitemaps/graph-els.xml", sitemapGraphEls)
	r.HandleFunc("/sitemaps/indecs-els.xml", sitemapIndexEls)
*/
