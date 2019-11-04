package graph

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type graphSitemap struct {
	Head *head.Head
	Tree *entry.Hold
}

func Check(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Title:   "Check - Graph",
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      s.Trees["graph"],
		Night:   head.NightMode(r),
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "graph-check", &graphSitemap{
		Head: head,
		Tree: s.Trees["graph"],
	})
	if err != nil {
		log.Println(err)
	}
	return
}
