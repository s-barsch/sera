package index

import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
)

type indexMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Recents entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	lang := head.Lang(r.Host)

	t := s.Trees["index"].Public[lang]

	head := &head.Head{
		Title:   "Index",
		Section: "index",
		Path:    path,
		Host:    r.Host,
		Entry:   t,
		Options: head.GetOptions(r),
	}
	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	// recents := s.Recents["index"]

	/*
	if s.Flags.Local {
		recents = s.Recents["index-private"]
	}
	*/

	err = s.ExecuteTemplate(w, "index-main", &indexMain{
		Head:    head,
		Tree:    t,
		//Recents: recents.Offset(0, 100).NoEmpty(lang),
	})
	if err != nil {
		log.Println(err)
	}
}
