package front

import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/server"
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
		Dark:    head.DarkColors(r),
		Large:   head.LargeType(r),
		NoLog:   head.LogMode(r),
	}
	err := head.Make()
	if err != nil {
		return
	}

	index := s.Recents["index"]
	graph := s.Recents["graph"]

	if s.Flags.Local {
		index = s.Recents["index-private"]
		graph = s.Recents["graph-private"]
	}

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:  head,
		Index: index.Offset(0, 100).NoEmpty(lang),
		Graph: graph.Offset(0, 100).NoEmpty(lang),
	})
	if err != nil {
		log.Println(err)
	}
}