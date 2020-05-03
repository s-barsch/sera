package about

import (
	"log"
	"net/http"
	//"stferal/go/entry"
	"stferal/go/entry/types/struct"
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

func ServeAbout(s *server.Server, w http.ResponseWriter, r *http.Request, struc *stru.Struct) {
	/*
	if perma := hold.Permalink(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}
	*/

	head := &head.Head{
		Title:   struc.Title(head.Lang(r.Host)),
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

	err = s.ExecuteTemplate(w, "about-main", &aboutStruct{
		Head:   head,
		Struct: struc,
	})
	if err != nil {
		log.Println(err)
	}
}
