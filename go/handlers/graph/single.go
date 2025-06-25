package graph

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type graphSingle struct {
	Meta  *meta.Meta
	Entry entry.Entry
	Prev  entry.Entry
	Next  entry.Entry
}

func ServeSingle(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	graph := s.Srv.Trees["graph"].Access(m.Auth.Subscriber)[m.Lang]
	e, err := graph.LookupEntryHash(m.Split.Hash)
	if err != nil {
		http.Redirect(w, r, "/graph", http.StatusMovedPermanently)
		return
	}

	perma := e.Perma(m.Lang)
	if m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	prev, next := getPrevNext(s.Srv.Recents["graph"].Access(m.Auth.Subscriber)[m.Lang], e)

	m.Title = graphEntryTitle(e, m.Lang)

	m.SetSection("graph")
	m.SetHreflang(e)

	/*
		schema, err := head.ElSchema()
		if err != nil {
			log.Println(err)
			return
		}
		head.Schema = schema
	*/

	err = s.Srv.ExecuteTemplate(w, "graph-single", &graphSingle{
		Meta:  m,
		Entry: e,
		Prev:  prev,
		Next:  next,
	})
	if err != nil {
		log.Println(err)
	}
}

func graphEntryDate(d time.Time, lang string) string {
	return fmt.Sprintf(d.Format("2 %v 2006"), tools.MonthLang(d, lang))
}

func graphEntryTitle(e entry.Entry, lang string) string {
	return fmt.Sprintf("%v - %v", e.Title(lang), graphEntryDate(e.Date(), lang))
}

func getPrevNext(es entry.Entries, single entry.Entry) (prev, next entry.Entry) {
	id := single.Id()
	for i, e := range es {
		if e.Id() == id {
			if i > 0 {
				next = es[i-1]
			}

			if i+1 < len(es) {
				prev = es[i+1]
			}
			return
		}
	}

	return
}
