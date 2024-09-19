package cache

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
)

/*
type cacheYear struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}
*/

func Year(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, p *paths.Path) {
	cache := s.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]

	id, err := getYearId(p.Slug)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	t, err := cache.LookupTree(id)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	m.Title = tools.Title(fmt.Sprintf("%v - %v", t.Date().Format("2006"), tools.KineName[m.Lang]))
	m.Section = "cache"
	// TODO:
	//m.Desc = s.Vars.Lang("cache-desc", m.Lang)

	err = m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := t.TraverseEntriesReverse()

	err = s.ExecuteTemplate(w, "cache-year", &cacheMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}

func getYearId(year string) (int64, error) {
	t, err := time.Parse("2006", year)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
