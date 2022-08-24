package front

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/entry/types/tree"
)

type frontMain struct {
	Head     *head.Head
	Index    entry.Entries
	Graph    entry.Entries
	Kine     entry.Entries
	Log      entry.Entries
	Months   tree.Trees
	Featured entry.Entry
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	lang := head.Lang(r.Host)
	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		Entry:   nil,
		Desc:    s.Vars.Lang("site", lang),
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		return
	}

	indecs := s.Recents["indecs"].Access(a.Subscriber)[lang]
	graph := s.Recents["graph"].Access(a.Subscriber)[lang]
	kine := s.Recents["kine"].Access(a.Subscriber)[lang]

	months := s.Trees["graph"].Access(a.Subscriber)[lang].TraverseTrees()
	newmonths := []*tree.Tree{}

	for _, m := range months {
		if m.Info()["release"] != "" {
			newmonths = append(newmonths, m)
		}
	}

	months = newmonths

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:  head,
		Index: indecs.Limit(s.Vars.FrontSettings.Index),
		Graph: graph.Limit(s.Vars.FrontSettings.Graph),
		Kine:  kine.Limit(10),
		Months: months,
		Log:   s.Recents["log"].Access(true)["de"].Limit(s.Vars.FrontSettings.Log),
	})
	if err != nil {
		log.Println(err)
	}
}

/*
	e, err := s.Trees["graph"].Access(a.Subscriber)[lang].LookupEntryHash(s.Vars.FrontSettings.Featured)
	if err != nil {
		s.Log.Println(err)
	}
*/
//Featured: e,
