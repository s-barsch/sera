package graph

/*
import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/server/head"
	"sacer/go/server"
)

type graphSitemap struct {
	Head *head.Head
	Tree *entry.Hold
}

func Check(s *server.Server, w http.ResponseWriter, r *http.Request, a *users.Auth) {
	head := &head.Head{
		Title:   "Check - Graph",
		Section: "graph",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      s.Trees["graph"],
		Dark:    head.DarkColors(r),
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "graph-check", &graphSitemap{
		Head: head,
		Tree: s.Trees["graph"],
	})
	if err != nil {
		log.Println(err)
	}
	return
}
*/
