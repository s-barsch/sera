package graph

/*
import (
	"fmt"
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/paths"
	"stferal/go/server"
	"time"
)

func Year(s *server.Server, w http.ResponseWriter, r *http.Request, p *paths.Path) {
	tree := s.Trees["graph"]

	if s.Flags.Local {
		tree = s.Trees["graph-private"]
	}

	h, err := findYear(tree, p)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	if perma := h.Permalink(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	// month
	if h.Depth() > 1 {
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
		El:      h,
		Dark:    head.DarkColors(r),
		Large:   head.LargeType(r),
	}
	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "graph-main", &graphMain{
		Head: head,
		Hold: h,
		Els:  serializeMonths(h),
		Prev: prev,
		Next: next,
	})
	if err != nil {
		log.Println(err)
	}
}

func findYear(tree *entry.Hold, p *paths.Path) (*entry.Hold, error) {
	timestamp, err := getTimestamp(p)
	if err != nil {
		return nil, err
	}

	eh, err := tree.Lookup(timestamp)
	if err != nil {
		return nil, err
	}
	h, ok := eh.(*entry.Hold)
	if !ok {
		return nil, fmt.Errorf("findYear: Could not convert interface to type Hold.")
	}
	return h, nil
}

func serializeMonths(h *entry.Hold) entry.Els {
	els := entry.Els{}
	for _, month := range h.Holds {
		for _, e := range month.Els {
			els = append(els, e)
		}
	}
	return els
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
	return t.Format(entry.Timestamp), nil
}

func yearTitle(h *entry.Hold, lang string) string {
	title := h.Info.Title(lang)
	if h.Depth() == 1 {
		title += " - Graph"
	}
	if h.Depth() == 2 {
		title += " " + h.Date.Format("2006")
	}
	return title
}
*/
