package graph

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/server/paths"
	"g.rg-s.com/sera/go/viewer"
)

type Viewer interface {
	View(w http.ResponseWriter, r *http.Request, m *meta.Meta)
}

type Graph struct {
	Server *s.Server
}

type Request struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Meta           *meta.Meta
	Path           Path
}

type Path struct {
}
type monthPage struct {
	Meta *meta.Meta
	Tree *tree.Tree
	Prev *tree.Tree
	Next *tree.Tree
}

func MonthPage(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getMonthId(m.Split)
		if err != nil {
			http.NotFound(w, r)
			log.Println(err)
			return
		}

		graph := s.Srv.Store.Trees["graph"].Access(m.Auth.Subscriber)[m.Lang]
		t, err := graph.LookupTree(id)
		if err != nil {
			http.NotFound(w, r)
			log.Println(err)
			return
		}

		if perma := t.Perma(m.Lang); m.Path != perma {
			http.Redirect(w, r, perma, http.StatusMovedPermanently)
			return
		}

		prev, next := prevNext(t)

		m.Title = monthTitle(t, m.Lang)
		m.Desc = metaDescription(t.Date(), m.Lang)

		m.SetSection("graph")
		m.SetHreflang(t)

		err = s.Srv.ExecuteTemplate(w, "graph-month", &monthPage{
			Meta: m,
			Tree: t,
			Prev: prev,
			Next: next,
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func metaDescription(d time.Time, lang string) string {
	if lang == "de" {
		return fmt.Sprintf("Monat %v %v des Graph von S. Feral.", tools.MonthLang(d, lang), d.Format("2006"))
	}
	return fmt.Sprintf("Month %v %v of Graph by S. Feral.", d.Format("January"), d.Format("2006"))
}

func monthTitle(t *tree.Tree, lang string) string {
	return fmt.Sprintf("%v %v - Graph", t.Title(lang), t.Date().Format("2006"))
}

func getMonthId(p *paths.Split) (int64, error) {
	if len(p.Chain) != 3 {
		return 0, fmt.Errorf("could not parse month id: %v", p.Raw)
	}

	slug := p.Slug
	if paths.IsMergedMonths(p.Slug) {
		slug = slug[:2]
	}

	t, err := time.Parse("2006-01", fmt.Sprintf("%v-%v", p.Chain[2], slug))
	if err != nil {
		return 0, err
	}
	// Years start on second 00, months on 01, days on 02. Hence, add a second.
	return t.Add(time.Second).Unix(), nil
}

func prevNext(t *tree.Tree) (prev, next *tree.Tree) {
	year, ok := t.Parent().(*tree.Tree)
	if !ok {
		return
	}
	for i, child := range year.Trees {
		if child.Id() == t.Id() {
			if i > 0 {
				prev = year.Trees[i-1]
			}
			if i+1 < len(year.Trees) {
				next = year.Trees[i+1]
			}
			if i == 0 {
				prev = prevYearLastMonth(year)
			}
			if i+1 == len(year.Trees) && i != 0 {
				next = nextYearFirstMonth(year)
			}
		}
	}
	return
}

func nextYearFirstMonth(year *tree.Tree) *tree.Tree {
	graph, ok := year.Parent().(*tree.Tree)
	if !ok {
		return nil
	}
	for i, child := range graph.Trees {
		if child.Id() == year.Id() {
			if i+1 < len(graph.Trees) {
				next := graph.Trees[i+1]
				if len(next.Trees) > 0 {
					return next.Trees[0]
				}
			}
		}
	}
	return nil
}

func prevYearLastMonth(year *tree.Tree) *tree.Tree {
	graph, ok := year.Parent().(*tree.Tree)
	if !ok {
		return nil
	}
	for i, child := range graph.Trees {
		if child.Id() == year.Id() {
			if i < 0 {
				prev := graph.Trees[i-1]
				if l := len(prev.Trees); l > 0 {
					return prev.Trees[l-1]
				}
			}
		}
	}
	return nil
}
