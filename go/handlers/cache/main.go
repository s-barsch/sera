package cache

import (
	//"fmt"
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type cacheMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	t := s.Store.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = "Cache"
	m.Desc = t.Info().Field("description", m.Lang)

	m.SetSection("cache")
	m.SetHreflang(t)

	entries := s.Store.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang].Limit(10)

	err := s.Store.ExecuteTemplate(w, "cache-main", &cacheMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}
