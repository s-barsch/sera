package graph

import (
	"fmt"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func MainRedirect(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	path := fmt.Sprintf("/%v/graph/2021/04", m.Lang)
	http.Redirect(w, r, path, http.StatusTemporaryRedirect)
}

func YearRedirect(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	t, err := findYear(m.Split.Slug)
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

func findYear(str string) (*tree.Tree, error) {
	graph := s.Store.Trees["graph"].Access(true)["de"]

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
