package kine

import (
	//"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/users"
	"sacer/go/server/head"
	"strings"
)

type kineMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	lang := head.Lang(r.Host)
	t := s.Trees["kine"].Access(a.Subscriber)[lang]
	head := &head.Head{
		Title:   strings.Title(tools.KineName[lang]),
		Section: "kine",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   t,
		Auth:    a,
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := s.Recents["kine"].Access(a.Subscriber)[lang]

	err = s.ExecuteTemplate(w, "kine-main", &kineMain{
		Head:    head,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}
