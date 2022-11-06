package graph

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
	"time"
	"fmt"
)

type monthPage struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Prev    *tree.Tree
	Next    *tree.Tree
}

func MonthPage(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, p *paths.Path) {
	graph := s.Trees["graph"].Access(m.Auth.Subscriber)[m.Lang]

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

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	m.Title = monthTitle(t, m.Lang)
	m.Section = "graph"

	err = m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "graph-month", &monthPage{
		Meta:	m,
		Tree:   t,
	})
	if err != nil {
		log.Println(err)
	}
}

func monthTitle(t *tree.Tree, lang string) string {
	return fmt.Sprintf("%v %v - Graph", t.Title(lang), t.Date().Format("2006"))
}

func getMonthId(p *paths.Path) (int64, error) {
	if len(p.Chain) != 3 {
		return 0, fmt.Errorf("Could not parse month id: %v", p.Raw)
	}
	t, err := time.Parse("2006-01", fmt.Sprintf("%v-%v", p.Chain[2], p.Slug))
	if err != nil {
		return 0, err
	}
	// Years start on second 00, months on 01, days on 02. Hence, add a second.
	return t.Add(time.Second).Unix(), nil
}


