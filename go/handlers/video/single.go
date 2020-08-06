package video

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
}

func ServeSingle(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	lang := head.Lang(r.Host)
	graph := s.Trees["video"].Local(s.Flags.Local)[lang]
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

	head := &head.Head{
		Title:   graphEntryTitle(e, head.Lang(r.Host)),
		Section: "video",
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

	err = s.ExecuteTemplate(w, "video-single", &graphSingle{
		Head:   head,
		Entry:  e,
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
