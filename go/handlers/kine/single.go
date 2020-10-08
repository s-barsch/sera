package kine

import (
	"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"strings"
	"time"
)

type graphSingle struct {
	Head  *head.Head
	Entry entry.Entry
}

func ServeSingle(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	lang := head.Lang(r.Host)
	graph := s.Trees["kine"].Access(a.Subscriber)[lang]
	e, err := graph.LookupEntryHash(p.Hash)
	if err != nil {
		http.Redirect(w, r, "/kine", 301)
		return
	}

	perma := e.Perma(lang)
	if r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   getTitle(e, head.Lang(r.Host)),
		Section: "kine",
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

	err = s.ExecuteTemplate(w, "kine-single", &graphSingle{
		Head:  head,
		Entry: e,
	})
	if err != nil {
		log.Println(err)
	}
}

func getDate(d time.Time, lang string) string {
	return fmt.Sprintf(d.Format("02 %v"), tools.Abbr(tools.MonthLang(d, lang)))
}

func getTitle(e entry.Entry, lang string) string {
	return fmt.Sprintf("%v - %v - %v", getDate(e.Date(), lang), e.Title(lang), strings.Title(tools.KineName[lang]))
}
