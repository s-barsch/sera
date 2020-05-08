package about

import (
	"log"
	"net/http"
	//"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"stferal/go/head"
	"stferal/go/server"
	/*
		"strings"
		"path/filepath"
		"stferal/go/entry"
		"stferal/go/head"
	*/)

type aboutTree struct {
	Head *head.Head
	Tree *tree.Tree
}

func ServeAbout(s *server.Server, w http.ResponseWriter, r *http.Request, t *tree.Tree) {
	lang := head.Lang(r.Host)
	if perma := t.Perma(lang); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   t.Title(lang),
		Section: "about",
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

	err = s.ExecuteTemplate(w, "about-main", &aboutTree{
		Head: head,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}
