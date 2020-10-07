package front

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/head"
	"sacer/go/server"
	"sacer/go/server/auth"
)

type frontMain struct {
	Head     *head.Head
	Index    entry.Entries
	Graph    entry.Entries
	Kine     entry.Entries
	Log      entry.Entries
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

	index := s.Recents["index"][lang]
	graph := s.Recents["graph"][lang]
	kine := s.Recents["kine"][lang]

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:     head,
		Index:    index.Limit(s.Vars.FrontSettings.Index),
		Graph:    graph.Limit(s.Vars.FrontSettings.Graph),
		Kine:     kine.Limit(10),
		Log:      s.Recents["log"]["de"],
	})
	if err != nil {
		log.Println(err)
	}
}

	/*
	e, err := s.Trees["graph"][lang].LookupEntryHash(s.Vars.FrontSettings.Featured)
	if err != nil {
		s.Log.Println(err)
	}
	*/
		//Featured: e,
