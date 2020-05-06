package graph

import (
	//"fmt"
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
	graph := s.Trees["graph"]

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

	if perma := tree.Perma(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	// month
	if tree.Level() >= 2 {
		http.NotFound(w, r)
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
		Title:   "not yet",//yearTitle(t, head.Lang(r.Host)),
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
	id, err := getId(p)
	if err != nil {
		return nil, err
	}

	return graph.LookupTree(id)
}

func getId(path *paths.Path) (int64, error) {
	t, err := getTime(path.Slug, firstEl(path.Parents))
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

func getTime(number, parent string) (time.Time, error) {
	// year
	if len(number) > 2 {
		return time.Parse("2006", number)
	}

	// month
	t, err := time.Parse("200601", parent+number)
	if err != nil {
		return t, err
	}
	if t.Month() == 1 {
		t = t.Add(time.Second)
	}

	return t, nil
}


/*
func yearTitle(h *entry.Hold, lang string) string {
	title := h.Info.Title(lang)
	if h.Depth() == 1 {
		title += " - Graph"
	}
	if h.Depth() == 2 {
		title += " " + h.Date.Format("2006")
	}
	return title
}
*/

func serializeMonths(tree *tree.Tree) entry.Entries {
	es := entry.Entries{}
	for _, month := range tree.Trees {
		for _, e := range month.Entries() {
			es = append(es, e)
		}
	}
	return es
}


