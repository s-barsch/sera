package graph

/*
import (
	"log"
	"net/http"
	"g.rg-s.com/sferal/go/entry"
	"g.rg-s.com/sferal/go/server/meta"
	"g.rg-s.com/sferal/go/server"
)

type graphSitemap struct {
	Head *meta.Meta
	Tree *entry.Hold
}

func Check(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &meta.Meta{
		Title:   "Check - Graph",
		Section: "graph",
		Path:    r.URL.Path,
		El:      s.Trees["graph"],
		Dark:    head.DarkColors(r),
	}
	err := head.Process(r)
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
