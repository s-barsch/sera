package index

import (
	"fmt"
	"log"
	"net/http"
	"stferal/go/entry/types/tree"
	"stferal/go/head"
	"stferal/go/server"
)

type indexPage struct {
	Head *head.Head
	Tree *tree.Tree
}

func IndexPage(s *server.Server, w http.ResponseWriter, r *http.Request, t *tree.Tree) {
	lang := head.Lang(r.Host)

	if perma := t.Perma(lang); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   indexTitle(t, lang),
		Section: "index",
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

	err = s.ExecuteTemplate(w, "index-page", &indexPage{
		Head: head,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}

func indexTitle(t *tree.Tree, lang string) string {
	title := t.Title(lang)

	topicTitle := ""

	if topic := t.TopicPage(); topic != nil {
		topicTitle = fmt.Sprintf(" - %v ", topic.Title(lang))
	}

	c := t.Chain()
	if len(c) > 2 {
		mainCategory := c[1].Title(lang)
		title += topicTitle + " - " + mainCategory
	}

	return title
}
