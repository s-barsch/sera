package auth

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server/users"
	"sacer/go/server/head"
	"sacer/go/server"
	"strings"
)

type extraHold struct {
	Head *head.Head
	Tree *tree.Tree
}


func SysPage(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	lang := head.Lang(r.Host)
	extra := s.Trees["extra"].Access(a.Subscriber)[lang]

	items := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	t, err := extra.SearchTree(items[len(items)-1], head.Lang(r.Host))
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	if perma := t.Perma(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}


	head := &head.Head{
		Title:   t.Title(lang),
		Section: "extra",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   t,
		Options: head.GetOptions(r),
	}
	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, t.Slug("en") + "-extra", &extraHold{
		Head: head,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}

