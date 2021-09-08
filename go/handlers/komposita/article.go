package komposita

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
)

type kompositaArticle struct {
	Head  *head.Head
	Entry entry.Entry
}

func Article(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth, p *paths.Path) {
	lang := head.Lang(r.Host)
	komposita := s.Trees["komposita"].Access(a.Subscriber)[lang]
	e, err := komposita.LookupEntryHash(p.Hash)
	if err != nil {
		http.NotFound(w, r)
		//http.Redirect(w, r, "/index", 301)
		return
	}

	perma := e.Perma(lang)
	if r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	head := &head.Head{
		Title:   e.Title(lang),
		Section: "komposita",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   e,
		Options: head.GetOptions(r),
	}

	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "komposita-article", &kompositaArticle{
		Head:  head,
		Entry: e,
	})
	if err != nil {
		log.Println(err)
	}
}
