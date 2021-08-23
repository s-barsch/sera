package indecs

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
	"sacer/go/server/paths"
)

type indecsMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Recents entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	lang := head.Lang(r.Host)

	t := s.Trees["indecs"].Access(a.Subscriber)[lang]

	head := &head.Head{
		Title:   "Indecs",
		Section: "indecs",
		Path:    path,
		Host:    r.Host,
		Entry:   t,
		Options: head.GetOptions(r),
	}
	err = head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	recents := s.Recents["indecs"].Access(a.Subscriber)[lang]

	err = s.ExecuteTemplate(w, "indecs-main", &indecsMain{
		Head:    head,
		Tree:    t,
		Recents: recents.Offset(0, 100),
	})
	if err != nil {
		log.Println(err)
	}
}
