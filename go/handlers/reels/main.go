package reels

import (
	//"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/meta"
	"strings"
)

type reelsMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	t := s.Trees["reels"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = strings.Title(tools.KineName[m.Lang])
	m.Section = "reels"
	m.Desc = t.Info().Field("description", m.Lang)

	err := m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := s.Recents["reels"].Access(m.Auth.Subscriber)[m.Lang].Limit(10)

	err = s.ExecuteTemplate(w, "reels-main", &reelsMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}
