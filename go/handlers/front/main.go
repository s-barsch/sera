package front

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func Rewrites(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
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

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	m.Title = ""
	m.Section = "home"
	m.Desc = s.Vars.Lang("site", m.Lang)

	err := m.Process(nil)
	if err != nil {
		return
	}

	//indecs := s.Recents["indecs"].Access(m.Auth.Subscriber)[m.Lang]
	graph := s.Recents["graph"].Access(m.Auth.Subscriber)[m.Lang]
	cache := s.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang]

	months := s.Trees["graph"].Access(m.Auth.Subscriber)[m.Lang].TraverseTrees()
	newmonths := []*tree.Tree{}

	for _, m := range months {
		if m.Info()["release"] != "" {
			newmonths = append(newmonths, m)
		}
	}

	months = newmonths

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Meta: m,
		//Index:  indecs.Limit(s.Vars.FrontSettings.Index),
		Graph:  graph.Limit(s.Vars.FrontSettings.Graph),
		Cache:  cache.Limit(10),
		Months: months,
		// Log:    s.Recents["log"].Access(true)["de"].Limit(s.Vars.FrontSettings.Log),
	})
	if err != nil {
		log.Println(err)
	}
}

/*
	e, err := s.Trees["graph"].Access(a.Subscriber)[m.Lang].LookupEntryHash(s.Vars.FrontSettings.Featured)
	if err != nil {
		s.Log.Println(err)
	}
*/
//Featured: e,
