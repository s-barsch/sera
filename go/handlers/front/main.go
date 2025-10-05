package front

import (
	"log"
	"net/http"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/types/tree"
	"g.rg-s.com/sacer/go/server/meta"
	"g.rg-s.com/sacer/go/viewer"
)

func Rewrites(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

type frontMain struct {
	Meta     *meta.Meta
	Index    entry.Entries
	Graph    entry.Entries
	Cache    entry.Entries
	Log      entry.Entries
	Months   tree.Trees
	Featured entry.Entry
}

func Main(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Title = ""
		m.Desc = v.Engine.Vars.Lang("site", m.Lang)

		m.SetSection("home")
		m.SetHreflang(nil)

		err := v.Engine.ExecuteTemplate(w, "front", &frontMain{
			Meta: m,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
