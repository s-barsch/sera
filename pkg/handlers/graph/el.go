package graph

import (
	"fmt"
	"log"
	"net/http"
	"st/pkg/el"
	"st/pkg/head"
	"st/pkg/paths"
	"st/pkg/server"
)

type graphEl struct {
	Head   *head.Head
	El     interface{}
	Parent *el.Hold
	Prev   interface{}
	Next   interface{}
}

func El(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	e, err := s.Trees["graph"].LookupAcronym(p.Acronym)
	if err != nil {
		http.Redirect(w, r, "/graph/", 301)
		return
	}

	perma := el.Permalink(e, head.Lang(r.Host))
	if r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	prev, next, err := getPrevNext(s.Recents["graph"], p.Acronym)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	parent, err := el.ElHold(e)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	title, err := elTitle(e, head.Lang(r.Host))
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	head := &head.Head{
		Title:   title,
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      e,
	}
	err = head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	schema, err := head.ElSchema()
	if err != nil {
		s.Log.Println(err)
		return
	}
	head.Schema = schema

	err = s.ExecuteTemplate(w, "graph-el", &graphEl{
		Head:   head,
		El:     e,
		Parent: parent,
		Prev:   prev,
		Next:   next,
	})
	if err != nil {
		log.Println(err)
	}
}

func titleDate(e interface{}, lang string) (string, error) {
	d, err := el.DateSafe(e)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(d.Format("2 %v 2006"), el.MonthLang(d, lang)), nil
}

func elTitle(e interface{}, lang string) (string, error) {
	title := el.Title(e, lang)

	date, err := titleDate(e, lang)
	if err != nil {
		return title, err
	}

	return fmt.Sprintf("%v - %v", title, date), nil
}

func getPrevNext(els el.Els, acronym string) (prev, next interface{}, Err error) {
	id, err := el.DecodeAcronym(acronym)
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
