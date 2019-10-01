package main

import (
	"log"
	"net/http"
	"path/filepath"
	"st/el"
)

type emblemPage struct {
	Head *head
	Nav  nav

	Lang string

	Arg el.Arg

	Compare int
}

func serveIndexEl(w http.ResponseWriter, r *http.Request, page, acronym string) {
	//TODO: Eliminating second lookup. One is already in Route.
	e, err := lookupAcronymMulti(page, acronym)
	if err != nil {
		debug(err)
		http.NotFound(w, r)
		return
	}

	lang := lang(r.Host)

	if page == "about" && lang == "de" {
		page = "über"
	}

	perma := el.Permalink(e, lang)
	if r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	title := el.Title(e, lang)

	if title == "" {
		date, err := el.DateSafe(e)
		if err == nil {
			title = el.EncodeAcronym(date)
		}
	}

	f, err := el.ElFileSafe(e)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	title += " - " + f.Hold.Info.Title(lang)

	head := newHead(title, r.URL.Path, e, lang)

	schema, err := newElSchema(e, title, lang)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	head.Schema = schema

	head.Description = schema.Description

	err = srv.executeTemplate(w, "index-el-page", &emblemPage{
		Head: head,
		Nav:  buildNav(page, lang),

		Lang:    lang,
		Compare: -1,

		Arg: el.Arg{
			El:   e,
			Lang: lang,
		},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func getActivePage(path, lang string) string {
	switch filepath.Base(path) {
	case "ueber", "einfluesse":
		if lang == "de" {
			return "über"
		} else {
			return "about"
		}
	}
	return "index"
}
