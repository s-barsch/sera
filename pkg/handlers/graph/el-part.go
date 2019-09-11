package graph

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"stferal/pkg/el"
	"stferal/pkg/head"
	"stferal/pkg/paths"
	"stferal/pkg/server"
)

type graphPart struct {
	Title  string            `json:"title"`
	URL    string            `json:"url"`
	Html   string            `json:"html"`
	Prev   string            `json:"prev"`
	Next   string            `json:"next"`
	Parent string            `json:"parent"`
	Langs  map[string]string `json:"langs"`
}

func ElPart(s *server.Server, w http.ResponseWriter, r *http.Request) {
	p := paths.Split(r.URL.Path)

	e, err := s.Trees["graph"].LookupAcronym(p.Acronym)
	if err != nil {
		http.NotFound(w, r)
		s.Debug(err)
		return
	}

	lang := head.Lang(r.Host)

	prev, next, err := getPrevNext(s.Recents["graph"], p.Acronym)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	parent, err := el.ElHold(e)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	title, err := elTitle(e, lang)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	head := &head.Head{
		Title:   title,
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      s.Trees["graph"],
	}
	err = head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	graphEl := &graphEl{
		Head:   head,
		El:     e,
		Parent: parent,
		Prev:   prev,
		Next:   next,
	}

	html := &bytes.Buffer{}
	err = s.ExecuteTemplate(html, "graph-el-part", graphEl)
	if err != nil {
		http.Error(w, "Internal Error", 500)
		log.Println(err)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(&graphPart{
		Title:  graphEl.Head.PageTitle(),
		URL:    el.Permalink(e, lang),
		Html:   html.String(),
		Prev:   el.Permalink(prev, lang),
		Next:   el.Permalink(next, lang),
		Parent: el.Permalink(graphEl.Parent, lang),
		// TODO: sketchy.
		Langs: map[string]string{
			"de": head.Langs.Hreflang("de").Href,
			"en": head.Langs.Hreflang("en").Href,
		},
	})

	if err != nil {
		http.Error(w, "Internal error.", 500)
		log.Println(err)
		return
	}
}
