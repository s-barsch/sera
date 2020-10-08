package about

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/head"
)

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

	err = s.ExecuteTemplate(w, aboutTemplate(t.Level()), &aboutTree{
		Head: head,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}

func aboutTemplate(level int) string {
	if level == 0 {
		return "about-main"
	}
	return "about-page"
}
