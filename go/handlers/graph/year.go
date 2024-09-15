package graph

import (
	"fmt"
	"net/http"
	"time"

	"g.rg-s.com/sacer/go/entry/types/tree"
	"g.rg-s.com/sacer/go/server"
	"g.rg-s.com/sacer/go/server/meta"
	"g.rg-s.com/sacer/go/server/paths"
)

func MainRedirect(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path := fmt.Sprintf("/%v/graph/2021/04", m.Lang)
	http.Redirect(w, r, path, http.StatusTemporaryRedirect)
}

/*
func MainRedirect(s *server.Server, w http.ResponseWriter, r *http.Request) {
	graph := s.Trees["graph"].Access(true)["de"]
	if len(graph.Trees) < 1 {
		http.Error(w, "graph not found", 500)
		return
	}
	http.Redirect(w, r, graph.Trees[len(graph.Trees)-1].Perma("de"), 307)
}
*/

func YearRedirect(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, p *paths.Path) {
	t, err := findYear(s, p.Slug)
	if err != nil {
		http.Redirect(w, r, "/graph", http.StatusTemporaryRedirect)
		return
	}
	if len(t.Trees) < 1 {
		http.Error(w, "year has no months", 404)
		return
	}
	http.Redirect(w, r, t.Trees[0].Perma(m.Lang), http.StatusTemporaryRedirect)
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
