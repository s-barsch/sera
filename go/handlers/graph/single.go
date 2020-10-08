package graph

import (
	"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/head"
	"sacer/go/paths"
	"sacer/go/server"
	"sacer/go/server/auth"
	"time"
)

type graphSingle struct {
	Head  *head.Head
	Entry entry.Entry
	Prev  entry.Entry
	Next  entry.Entry
}

func ServeSingle(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	lang := head.Lang(r.Host)
	graph := s.Trees["graph"].Access(a.Subscriber)[lang]
	e, err := graph.LookupEntryHash(p.Hash)
	if err != nil {
		http.Redirect(w, r, "/graph", 301)
		return
	}

	perma := e.Perma(lang)
	if r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	prev, next := getPrevNext(s.Recents["graph"].Access(a.Subscriber)[lang], e)

	head := &head.Head{
		Title:   graphEntryTitle(e, head.Lang(r.Host)),
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   e,
		Options: head.GetOptions(r),
	}

	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	/*
		schema, err := head.ElSchema()
		if err != nil {
			s.Log.Println(err)
			return
		}
		head.Schema = schema
	*/

	err = s.ExecuteTemplate(w, "graph-single", &graphSingle{
		Head:  head,
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
