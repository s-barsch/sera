package graph

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/tools"
	"g.rg-s.com/sacer/go/requests/meta"
	"g.rg-s.com/sacer/go/viewer"
)

type graphSingle struct {
	Meta  *meta.Meta
	Entry entry.Entry
	Prev  entry.Entry
	Next  entry.Entry
}

func ServeSingle(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e, err := v.Store.Graph().LookupEntryHash(m.Split.Hash)
		if err != nil {
			http.Redirect(w, r, "/graph", http.StatusMovedPermanently)
			return
		}

		perma := e.Perma(m.Lang)
		if m.Path != perma {
			http.Redirect(w, r, perma, http.StatusMovedPermanently)
			return
		}

		prev, next := getPrevNext(v.Store.GraphFlat(), e)

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

		err = v.Engine.ExecuteTemplate(w, "graph-single", &graphSingle{
			Meta:  m,
			Entry: e,
			Prev:  prev,
			Next:  next,
		})
		if err != nil {
			log.Println(err)
		}
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
