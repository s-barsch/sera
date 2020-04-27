package extra

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/paths"
	"stferal/pkg/server"
	"strings"
)

type extraHold struct {
	Head *head.Head
	Hold *entry.Hold
}

func Route(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	items := strings.Split(strings.Trim(path, "/"), "/")

	h, err := s.Trees["extra"].Search(items[len(items)-1], head.Lang(r.Host))
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	Extra(s, w, r, h)
}

func Extra(s *server.Server, w http.ResponseWriter, r *http.Request, h *entry.Hold) {
	if perma := h.Permalink(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   h.Info.Title(head.Lang(r.Host)),
		Section: "extra",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      h,
		Dark:   head.DarkColors(r),
		Large:   head.LargeType(r),
		NoLog:   head.LogMode(r),
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "extra-page", &extraHold{
		Head: head,
		Hold: h,
	})
	if err != nil {
		log.Println(err)
	}
}
