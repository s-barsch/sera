package graph

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"time"
)

type yearPage struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
	Prev    *tree.Tree
	Next    *tree.Tree
}

func YearPage(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	lang := head.Lang(r.Host)

	graph := s.Trees["graph"].Access(a.Subscriber)[lang]

	id, err := getYearId(p.Slug)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	t, err := graph.LookupTree(id)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	if perma := t.Perma(lang); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	prev, next := yearSiblings(t)

	head := &head.Head{
		Title:   yearTitle(t, lang),
		Section: "graph",
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

	err = s.ExecuteTemplate(w, "graph-year", &yearPage{
		Head:    head,
		Tree:    t,
		Entries: serializeMonths(t),
		Prev:    prev,
		Next:    next,
	})
	if err != nil {
		log.Println(err)
	}
}

func firstEl(slugs []string) string {
	if len(slugs) < 1 {
		return ""
	}
	return slugs[0]
}

func yearTitle(t *tree.Tree, lang string) string {
	title := t.Title(lang)
	if t.Level() == 1 {
		title += " - Graph"
	}
	if t.Level() == 2 {
		title += " " + t.Date().Format("2006")
	}
	return title
}

func serializeMonths(tree *tree.Tree) entry.Entries {
	es := entry.Entries{}
	for _, month := range tree.Trees {
		for _, e := range month.Entries() {
			es = append(es, e)
		}
	}
	return es
}
