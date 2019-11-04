package about

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
	/*
		"strings"
		"path/filepath"
		"stferal/pkg/entry"
		"stferal/pkg/head"
	*/)

type aboutHold struct {
	Head *head.Head
	Hold *entry.Hold
}

var aboutName = map[string]string{
	"de": "über",
	"en": "about",
}

func Hold(s *server.Server, w http.ResponseWriter, r *http.Request, hold *entry.Hold) {
	if perma := hold.Permalink(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   hold.Info.Title(head.Lang(r.Host)),
		Section: "about",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      hold,
		Night:   head.NightMode(r),
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "about", &aboutHold{
		Head: head,
		Hold: hold,
	})
	if err != nil {
		log.Println(err)
	}
}
