package graph

import (
	//"fmt"
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"stferal/go/head"
	"stferal/go/server"
)

type graphMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
	//Prev *entry.Hold
	//Next *entry.Hold
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Title:   "Graph",
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   s.Trees["graph"],
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	/*
		prev, _, err := yearSiblings(lastHold(s.Trees["graph"]))
		if err != nil {
			s.Log.Println(err)
			return
		}
	*/

	entries := s.Recents["graph"]
	/*
		if s.Flags.Local {
			els = s.Recents["graph-private"]
		}
	*/

	err = s.ExecuteTemplate(w, "graph-main", &graphMain{
		Head:    head,
		Tree:    s.Trees["graph"],
		Entries: entries, //els.NoEmpty(head.Lang).Offset(0, 100),
		//Prev: prev,
	})
	if err != nil {
		log.Println(err)
	}
}

/*
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
*/
