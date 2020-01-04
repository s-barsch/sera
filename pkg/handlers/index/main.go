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
		Night:   head.NightMode(r),
		Large:   head.TypeMode(r),
		NoLog:   head.LogMode(r),
	}
	err = head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	recents := s.Recents["index"].Offset(0, 100)

	err = s.ExecuteTemplate(w, "index-hold", &indexMain{
		Head:    head,
		Hold:    s.Trees["index"],
		Recents: deleteEmpty(recents, lang),
	})
	if err != nil {
		log.Println(err)
	}
}

// Temporary workaround
func deleteEmpty(entries entry.Els, lang string) entry.Els {
	clean := entry.Els{}
	for _, e := range entries {
		if entry.Type(e) == "text" {
			if e.(*entry.Text).Text[lang] == "" {
				continue
			}
		}
		clean = append(clean, e)
	}
	return clean
}
