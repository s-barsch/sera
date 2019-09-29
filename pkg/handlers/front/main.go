package front

import (
	"log"
	"net/http"
	"stferal/pkg/el"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type frontMain struct {
	Head  *head.Head
	Index el.Els
	Graph el.Els
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Title:   "",
		Section: "home",
		Path:    "/",
		Host:    r.Host,
		El:      nil,
		Desc:    s.Vars.Lang("front", head.Lang(r.Host)),
		Dark:    head.DarkMode(r),
	}
	err := head.Make()
	if err != nil {
		return
	}

	err = s.ExecuteTemplate(w, "front", &frontMain{
		Head:  head,
		Index: s.Recents["index"],
		Graph: s.Recents["graph"],
	})
	if err != nil {
		log.Println(err)
	}
}
