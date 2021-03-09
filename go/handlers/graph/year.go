package graph

import (
	"net/http"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"sacer/go/entry/types/tree"
	"time"
)

func MainRedirect(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	graph := s.Trees["graph"].Access(true)["de"]
	if len(graph.Trees) < 1 {
		http.Error(w, "graph not found", 500)
		return
	}
	http.Redirect(w, r, graph.Trees[len(graph.Trees)-1].Perma("de"), 307)
}

func YearRedirect(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	t, err := findYear(s, p.Slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if len(t.Trees) < 1 {
		http.Error(w, "year has no months", 404)
		return
	}
	http.Redirect(w, r, t.Trees[0].Perma(head.Lang(r.Host)), 307)
	return
}

func findYear(s *server.Server, str string) (*tree.Tree, error) {
	graph := s.Trees["graph"].Access(true)["de"]

	id, err := getYearId(str)
	if err != nil {
		return nil, err
	}

	return graph.LookupTree(id)
}

func getYearId(year string) (int64, error) {
	t, err := time.Parse("2006", year)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}


