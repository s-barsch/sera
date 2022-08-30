package extra

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"sacer/go/server"
	"sacer/go/server/users"
	"strings"
)

type extraHold struct {
	Head *head.Head
	Tree *tree.Tree
}

func Route(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	items := strings.Split(strings.Trim(path, "/"), "/")

	lang := head.Lang(r.Host)
	extra := s.Trees["extra"].Access(a.Sub())[lang]

	t, err := extra.SearchTree(items[len(items)-1], head.Lang(r.Host))
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	Extra(s, w, r, a, t)
}

func Extra(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth, t *tree.Tree) {
	if perma := t.Perma(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	lang := head.Lang(r.Host)

	head := &head.Head{
		Title:   t.Title(lang),
		Section: "extra",
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

	err = s.ExecuteTemplate(w, "extra-page", &extraHold{
		Head: head,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}
