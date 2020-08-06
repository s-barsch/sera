package video

import (
	//"fmt"
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"stferal/go/head"
	"stferal/go/server"
)

type videoMain struct {
	Head    *head.Head
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)
	t := s.Trees["video"].Local(s.Flags.Local)[lang]
	head := &head.Head{
		Title:   "Video",
		Section: "video",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   t,
		Options: head.GetOptions(r),
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := s.Recents["video"].Local(s.Flags.Local)[lang]

	err = s.ExecuteTemplate(w, "video-main", &videoMain{
		Head:    head,
		Tree:    t,
		Entries: entries.Offset(0, 100),
	})
	if err != nil {
		log.Println(err)
	}
}

