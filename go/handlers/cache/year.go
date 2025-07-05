package cache

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

func Year(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getYearId(m.Split.Slug)
		if err != nil {
			http.NotFound(w, r)
			log.Println(err)
			return
		}

		t, err := v.Store.Cache().LookupTree(id)
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
		// TODO: m.Desc = s.Engine.Vars.Lang("cache-desc", m.Lang)
		m.SetSection("cache")
		m.SetHreflang(t)

		entries := t.Flatten()

		err = v.Engine.ExecuteTemplate(w, "cache-year", &cacheMain{
			Meta:    m,
			Tree:    t,
			Entries: entries,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func getYearId(year string) (int64, error) {
	t, err := time.Parse("2006", year)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
