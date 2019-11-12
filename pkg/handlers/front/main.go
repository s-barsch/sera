package front

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type frontMain struct {
	Head  *head.Head
	Index entry.Els
	Graph entry.Els
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)

	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		El:      nil,
		Desc:    s.Vars.Lang("site", head.Lang(r.Host)),
		Night:   head.NightMode(r),
		Large:   head.TypeMode(r),
	}
	err := head.Make()
	if err != nil {
		return
	}

	index := s.Recents["index"].Offset(0, 100)
	graph := s.Recents["graph"].Offset(0, 100)

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:  head,
		Index: deleteEmpty(index, lang),
		Graph: deleteEmpty(graph, lang),
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
