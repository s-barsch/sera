package graph

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"time"
	"fmt"
)

type monthPage struct {
	Head    *head.Head
	Tree    *tree.Tree
	Prev    *tree.Tree
	Next    *tree.Tree
}

func MonthPage(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	lang := head.Lang(r.Host)

	graph := s.Trees["graph"].Access(a.Subscriber)[lang]

	id, err := getMonthId(p)
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

	head := &head.Head{
		Title:   monthTitle(t, lang),
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

	err = s.ExecuteTemplate(w, "graph-month", &monthPage{
		Head:    head,
		Tree:    t,
	})
	if err != nil {
		log.Println(err)
	}
}

func monthTitle(t *tree.Tree, lang string) string {
	return fmt.Sprintf("%v %v", t.Title(lang), t.Date().Format("2006"))
}

func getMonthId(p *paths.Path) (int64, error) {
	if len(p.Chain) != 2 {
		return 0, fmt.Errorf("Could not parse month id: %v", p.Raw)
	}
	t, err := time.Parse("2006-01", fmt.Sprintf("%v-%v", p.Chain[1], p.Slug))
	if err != nil {
		return 0, err
	}
	// Years start on second 00, months on 01, days on 02. Hence, add a second.
	return t.Add(time.Second).Unix(), nil
}


