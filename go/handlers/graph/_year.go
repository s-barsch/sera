package graph

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
)

type yearPage struct {
	Head    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
	Prev    *tree.Tree
	Next    *tree.Tree
}

func YearPage(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
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

	if perma := t.Perma(lang); m.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	prev, next := yearSiblings(t)

	head := &meta.Meta{
		Title:   yearTitle(t, lang),
		Section: "graph",
		Path:    r.URL.Path,
		Entry:   t,
	}

	err = head.Process(r)
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
