package about

import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/server"
	/*
		"strings"
		"path/filepath"
		"stferal/go/entry"
		"stferal/go/head"
	*/)

type aboutStruct struct {
	Head   *head.Head
	Struct *stru.Struct
}

var aboutName = map[string]string{
	"de": "über",
	"en": "about",
}

func ServeStruct(s *server.Server, w http.ResponseWriter, r *http.Request, strct *stru.Struct) {
	/*
	if perma := hold.Permalink(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}
	*/
	log.Println("permalink unavailable")

	head := &head.Head{
		Title:   "About title missing",
		//Title:   hold.Info.Title(head.Lang(r.Host)),
		Section: "about",
		Path:    r.URL.Path,
		Host:    r.Host,
		// Entry:   hold,
		Options: head.GetOptions(r),
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "about", &aboutStruct{
		Head:   head,
		Struct: strct,
	})
	if err != nil {
		log.Println(err)
	}
}
