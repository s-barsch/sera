package extra

import (
	"log"
	"net/http"
	"stferal/go/entry/types/tree"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
	"strings"
)

type extraHold struct {
	Head *head.Head
	Tree *tree.Tree
}

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	items := strings.Split(strings.Trim(path, "/"), "/")

	lang := head.Lang(r.Host)
	extra := s.Trees["extra"].Local(s.Flags.Local)[lang]

	t, err := extra.SearchTree(items[len(items)-1], head.Lang(r.Host))
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	Extra(s, w, r, t)
}

func Extra(s *server.Server, w http.ResponseWriter, r *http.Request, t *tree.Tree) {
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