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
	Index entry.Entries
	Graph entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)

	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		//El:      nil,
		//Desc:    s.Vars.Lang("site", head.Lang(r.Host)),
		Options: head.GetOptions(r),
	}
	err := head.Process()
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
