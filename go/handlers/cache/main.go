package cache

import (
	//"fmt"
	"log"
	"net/http"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/types/tree"
	"g.rg-s.com/sacer/go/server/meta"
	"g.rg-s.com/sacer/go/viewer"
)

type cacheMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := v.Store.Cache()
		m.Title = "Cache"
		m.Desc = t.Info().Field("description", m.Lang)

		m.SetSection("cache")
		m.SetHreflang(t)

		err := v.Engine.ExecuteTemplate(w, "cache-main", &cacheMain{
			Meta:    m,
			Tree:    t,
			Entries: v.Store.CacheFlat().Limit(10),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
