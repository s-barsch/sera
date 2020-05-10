package graph

import (
	"fmt"
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
	"time"
	"stferal/go/entry/helper"
)

type graphSingle struct {
	Head   *head.Head
	Entry  entry.Entry
	Prev   entry.Entry
	Next   entry.Entry
}

func ServeSingle(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	lang := head.Lang(r.Host)
	graph := s.Trees["graph"].Local(s.Flags.Local)[lang]
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

	/*
	prev, next, err := getPrevNext(s.Recents["graph"], p.Acronym)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}
	*/

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
		Head:   head,
		Entry:  e,
		/*
		Parent: parent,
		Prev:   prev,
		Next:   next,
		*/
	})
	if err != nil {
		log.Println(err)
	}
}

func graphEntryDate(d time.Time, lang string) string {
	return fmt.Sprintf(d.Format("2 %v 2006"), helper.MonthLang(d, lang))
}

func graphEntryTitle(e entry.Entry, lang string) string {
	return fmt.Sprintf("%v - %v", e.Title(lang), graphEntryDate(e.Date(), lang))
}

/*
func getPrevNext(els entry.Els, acronym string) (prev, next interface{}, Err error) {
	id, err := entry.DecodeAcronym(acronym)
	if err != nil {
		Err = err
		return
	}
	i, err := els.LookupPosition(id)
	if err != nil {
		Err = err
		return
	}

	if i > 0 {
		prev = els[i-1]
	}

	if i+1 < len(els) {
		next = els[i+1]
	}
	return
}
*/
