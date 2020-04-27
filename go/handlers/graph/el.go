package graph

import (
	"fmt"
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
)

type graphEl struct {
	Head   *head.Head
	El     interface{}
	Parent *entry.Hold
	Prev   interface{}
	Next   interface{}
}

func El(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	e, err := s.Trees["graph"].LookupAcronym(p.Acronym)
	if err != nil {
		http.Redirect(w, r, "/graph/", 301)
		return
	}

	if entry.InfoSafe(e)["private"] == "true" {
		if !s.Flags.Local {
			http.Error(w, "403", 403)
			return
		}
	}

	perma := entry.Permalink(e, head.Lang(r.Host))
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

	parent, err := entry.ElHold(e)
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
		Dark:    head.DarkColors(r),
		Large:   head.LargeType(r),
		NoLog:   head.LogMode(r),
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
	d, err := entry.DateSafe(e)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(d.Format("2 %v 2006"), entry.MonthLang(d, lang)), nil
}

func elTitle(e interface{}, lang string) (string, error) {
	title := entry.Title(e, lang)

	date, err := titleDate(e, lang)
	if err != nil {
		return title, err
	}

	return fmt.Sprintf("%v - %v", title, date), nil
}

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
