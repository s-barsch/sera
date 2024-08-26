package sitemaps

import (
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
)

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	switch r.URL.Path {
	case "/sitemaps.xml":
		Index(s, w, r, m)
	case "/sitemaps/core.xml":
		Core(s, w, r, m)
	case "/sitemaps/trees.xml":
		Trees(s, w, r, m)
	case "/sitemaps/reelss.xml":
		Kines(s, w, r, m)
	case "/sitemaps/graph-entries.xml":
		GraphEntries(s, w, r, m)
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
