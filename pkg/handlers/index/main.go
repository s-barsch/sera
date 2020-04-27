package index

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/paths"
	"stferal/pkg/server"
)

type indexMain struct {
	Head    *head.Head
	Hold    *entry.Hold
	Recents entry.Els
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	lang := head.Lang(r.Host)

	head := &head.Head{
		Title:   "Index",
		Section: "index",
		Path:    path,
		Host:    r.Host,
		El:      s.Trees["index"],
		Dark:    head.DarkColors(r),
		Large:   head.LargeType(r),
		NoLog:   head.LogMode(r),
	}
	err = head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	recents := s.Recents["index"]

	if s.Flags.Local {
		recents = s.Recents["index-private"]
	}

	err = s.ExecuteTemplate(w, "index-hold", &indexMain{
		Head:    head,
		Hold:    s.Trees["index"],
		Recents: recents.Offset(0, 100).NoEmpty(lang),
	})
	if err != nil {
		log.Println(err)
	}
}
