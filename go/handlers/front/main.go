package front

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
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
		m.Desc = s.Srv.Engine.Vars.Lang("site", m.Lang)

		m.SetSection("home")
		m.SetHreflang(nil)

		//indecs := s.Store.Recents["indecs"].Access(m.Auth.Subscriber)[m.Lang]
		graph := s.Srv.Store.Recents["graph"].Access(m.Auth.Subscriber)[m.Lang]
		cache := s.Srv.Store.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang]

		months := s.Srv.Store.Trees["graph"].Access(m.Auth.Subscriber)[m.Lang].TraverseTrees()
		newmonths := []*tree.Tree{}

		for _, m := range months {
			if m.Info()["release"] != "" {
				newmonths = append(newmonths, m)
			}
		}

		months = newmonths

		err := s.Srv.ExecuteTemplate(w, "front", &frontMain{
			Meta: m,
			//Index:  indecs.Limit(s.Engine.Vars.FrontSettings.Index),
			Graph:  graph.Limit(s.Srv.Engine.Vars.FrontSettings.Graph),
			Cache:  cache.Limit(10),
			Months: months,
			// Log:    s.Store.Recents["log"].Access(true)["de"].Limit(s.Engine.Vars.FrontSettinglog),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
