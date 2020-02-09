package graph

import (
	"fmt"
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type graphMain struct {
	Head *head.Head
	Hold *entry.Hold
	Els  entry.Els
	Prev *entry.Hold
	Next *entry.Hold
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Title:   "Graph",
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      s.Trees["graph"],
		Night:   head.NightMode(r),
		Large:   head.TypeMode(r),
		NoLog:   head.LogMode(r),
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	prev, _, err := yearSiblings(lastHold(s.Trees["graph"]))
	if err != nil {
		s.Log.Println(err)
		return
	}

	els := s.Recents["graph"]
	if !s.Flags.Local {
		els = els.ExcludePrivate()
	}

	err = s.ExecuteTemplate(w, "graph-main", &graphMain{
		Head: head,
		Hold: s.Trees["graph"],
		Els:  els.Offset(0, 100),
		Prev: prev,
	})
	if err != nil {
		log.Println(err)
	}
}

func lastHold(hold *entry.Hold) *entry.Hold {
	if len(hold.Holds) < 1 {
		return nil
	}
	return hold.Holds.Reverse()[0]
}

func yearSiblings(h *entry.Hold) (prev, next *entry.Hold, err error) {
	if h == nil {
		err = fmt.Errorf("yearSiblings: Hold is nil.")
		return
	}
	if h.Mother == nil {
		err = fmt.Errorf("yearSiblings: Mother is nil.")
		return
	}

	for i, child := range h.Mother.Holds {
		if h.File.Id == child.File.Id {
			if i > 0 {
				prev = h.Mother.Holds[i-1]
			}

			if i+1 < len(h.Mother.Holds) {
				next = h.Mother.Holds[i+1]
			}
		}
	}
	return
}
