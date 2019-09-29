package graph

import (
	"fmt"
	"log"
	"net/http"
	"stferal/pkg/el"
	"stferal/pkg/head"
	"stferal/pkg/paths"
	"stferal/pkg/server"
	"time"
)

func Year(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	timestamp, err := getTimestamp(p)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}
	eh, err := s.Trees["graph"].Lookup(timestamp)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}
	h, ok := eh.(*el.Hold)
	if !ok {
		http.NotFound(w, r)
		return
	}

	prev, next, err := yearSiblings(h)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	head := &head.Head{
		Title:   yearTitle(h, head.Lang(r.Host)),
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      eh,
		Dark:    head.DarkMode(r),
	}
	err = head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "graph-main", &graphMain{
		Head: head,
		Hold: h,
		Els:  h.TraverseEls(),
		Prev: prev,
		Next: next,
	})
	if err != nil {
		log.Println(err)
	}
}

func getTime(p *paths.Path) (time.Time, error) {
	if len(p.Acronym) > 2 {
		return time.Parse("2006", p.Acronym)
	}
	if len(p.Holds) < 1 {
		return time.Time{}, fmt.Errorf("getTimestamp: Couldn’t determine date of %+v", p)
	}
	t, err := time.Parse("200601", p.Holds[0]+p.Acronym)
	if err != nil {
		return t, err
	}
	if t.Month() == 1 {
		t = t.Add(time.Second)
	}
	return t, nil
}

func getTimestamp(p *paths.Path) (string, error) {
	t, err := getTime(p)
	if err != nil {
		return "", err
	}
	return t.Format(el.Timestamp), nil
}

func yearTitle(h *el.Hold, lang string) string {
	title := h.Info.Title(lang)
	if h.Depth() == 1 {
		title += " - Graph"
	}
	if h.Depth() == 2 {
		title += " " + h.Date.Format("2006")
	}
	return title
}
