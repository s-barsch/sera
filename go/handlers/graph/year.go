package graph

import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
	"stferal/go/entry/types/tree"
	"time"
)

type yearPage struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
	//Prev *entry.Hold
	//Next *entry.Hold
}


func YearPage(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	lang := head.Lang(r.Host)
	graph := s.Trees["graph"].Public[lang]

	/*
	if s.Flags.Local {
		tree = s.Trees["graph-private"]
	}
	*/

	tree, err := findYearTree(graph, p)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}


	if perma := tree.Perma(lang); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	/*
	prev, next, err := yearSiblings(h)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}
	*/

	head := &head.Head{
		Title:   yearTitle(tree, lang),
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   tree,
		Options: head.GetOptions(r),
	}

	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "graph-year", &yearPage{
		Head:   head,
		Tree:    tree,
		Entries: serializeMonths(tree),
		/*
		Prev: prev,
		Next: next,
		*/
	})
	if err != nil {
		log.Println(err)
	}
}

func findYearTree(graph *tree.Tree, p *paths.Path) (*tree.Tree, error) {
	id, err := getId(p.Slug)
	if err != nil {
		return nil, err
	}

	return graph.LookupTree(id)
}

func getId(year string) (int64, error) {
	t, err := time.Parse("2006", year)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
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


