package kino

import (
	//"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"strings"
)

type kinoMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	lang := head.Lang(r.Host)
	t := s.Trees["kino"].Access(a.Subscriber)[lang]
	head := &head.Head{
		Title:   strings.Title(tools.KinoName[lang]),
		Section: "kino",
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

	entries := s.Recents["kino"].Access(a.Subscriber)[lang]

	err = s.ExecuteTemplate(w, "kino-main", &kinoMain{
		Head:    head,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}
