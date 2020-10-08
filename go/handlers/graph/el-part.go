package graph

/*
import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/server/head"
	"sacer/go/server/paths"
	"sacer/go/server"
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

func ElPart(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	p := paths.Split(r.URL.Path)

	e, err := s.Trees["graph"].LookupAcronym(p.Acronym)
	if err != nil {
		http.NotFound(w, r)
		s.Debug(err)
		return
	}

	if entry.InfoSafe(e)["private"] == "true" {
		if !s.Flags.Local {
			http.Error(w, "403", 403)
			return
		}
	}

	lang := head.Lang(r.Host)

	prev, next, err := getPrevNext(s.Recents["graph"], p.Acronym)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	parent, err := entry.ElHold(e)
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
	err = head.Process()
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
		URL:    entry.Permalink(e, lang),
		Html:   html.String(),
		Prev:   entry.Permalink(prev, lang),
		Next:   entry.Permalink(next, lang),
		Parent: entry.Permalink(graphEl.Parent, lang),
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
*/
