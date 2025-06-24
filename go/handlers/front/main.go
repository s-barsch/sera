package front

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func Rewrites(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
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

func Main(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	m.Title = ""
	m.Desc = s.Store.Vars.Lang("site", m.Lang)

	m.SetSection("home")
	m.SetHreflang(nil)

	//indecs := s.Recents["indecs"].Access(m.Auth.Subscriber)[m.Lang]
	graph := s.Store.Recents["graph"].Access(m.Auth.Subscriber)[m.Lang]
	cache := s.Store.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang]

	months := s.Store.Trees["graph"].Access(m.Auth.Subscriber)[m.Lang].TraverseTrees()
	newmonths := []*tree.Tree{}

	for _, m := range months {
		if m.Info()["release"] != "" {
			newmonths = append(newmonths, m)
		}
	}

	months = newmonths

	err := s.Store.ExecuteTemplate(w, "front", &frontMain{
		Meta: m,
		//Index:  indecs.Limit(s.Vars.FrontSettings.Index),
		Graph:  graph.Limit(s.Store.Vars.FrontSettings.Graph),
		Cache:  cache.Limit(10),
		Months: months,
		// Log:    s.Recents["log"].Access(true)["de"].Limit(s.Vars.FrontSettinglog),
	})
	if err != nil {
		log.Println(err)
	}
}
