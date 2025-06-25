package cache

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry/tools"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func Year(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	id, err := getYearId(m.Split.Slug)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}

	cache := s.Srv.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]
	t, err := cache.LookupTree(id)
	if err != nil {
		http.NotFound(w, r)
		log.Println(err)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	m.Title = tools.Title(fmt.Sprintf("%v - %v", t.Date().Format("2006"), "Cache"))
	// TODO: m.Desc = s.Vars.Lang("cache-desc", m.Lang)
	m.SetSection("cache")
	m.SetHreflang(t)

	entries := t.TraverseEntriesReverse()

	err = s.Srv.ExecuteTemplate(w, "cache-year", &cacheMain{
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
