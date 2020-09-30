package graph

import (
	//"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/types/tree"
	"sacer/go/head"
	"sacer/go/server"
)

type graphMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
	Prev    *tree.Tree
	//Next *entry.Hold
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)
	t := s.Trees["graph"][lang]
	head := &head.Head{
		Title:   "Graph",
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   t,
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	prev, _ := yearSiblings(lastTree(t))

	entries := s.Recents["graph"][lang]

	err = s.ExecuteTemplate(w, "graph-main", &graphMain{
		Head:    head,
		Tree:    t,
		Entries: entries.Offset(0, 100),
		Prev: prev,
	})
	if err != nil {
		log.Println(err)
	}
}

func lastTree(tree *tree.Tree) *tree.Tree {
	if len(tree.Trees) < 1 {
		return nil
	}
	return tree.Trees.Reverse()[0]
}

func yearSiblings(t *tree.Tree) (prev, next *tree.Tree) {
	if t == nil {
		return
	}
	if t.Parent() == nil {
		return
	}

	parentTree, ok := t.Parent().(*tree.Tree)
	if !ok {
		return
	}

	for i, child := range parentTree.Trees {
		if child.Id() == t.Id() {
			if i > 0 {
				prev = parentTree.Trees[i-1]
			}

			if i+1 < len(parentTree.Trees) {
				next = parentTree.Trees[i+1]
			}
		}
	}
	return
}
