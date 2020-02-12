package front

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type frontMain struct {
	Head  *head.Head
	Index entry.Els
	Graph entry.Els
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)

	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		El:      nil,
		Desc:    s.Vars.Lang("site", head.Lang(r.Host)),
		Night:   head.NightMode(r),
		Large:   head.TypeMode(r),
		NoLog:   head.LogMode(r),
	}
	err := head.Make()
	if err != nil {
		return
	}

	index := s.Recents["index"].Offset(0, 100)
	graph := s.Recents["graph"].Offset(0, 100)

	if !s.Flags.Local {
		graph = graph.ExcludePrivate()
		index = index.ExcludePrivate()
	}

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:  head,
		Index: index.NoEmpty(lang),
		Graph: graph.NoEmpty(lang),
	})
	if err != nil {
		log.Println(err)
	}
}
