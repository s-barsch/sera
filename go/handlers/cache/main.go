package cache

import (
	//"fmt"
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

type cacheMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := s.Srv.Store.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]

		m.Title = "Cache"
		m.Desc = t.Info().Field("description", m.Lang)

		m.SetSection("cache")
		m.SetHreflang(t)

		entries := s.Srv.Store.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang].Limit(10)

		err := s.Srv.ExecuteTemplate(w, "cache-main", &cacheMain{
			Meta:    m,
			Tree:    t,
			Entries: entries,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
