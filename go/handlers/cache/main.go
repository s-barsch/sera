package cache

import (
	//"fmt"
	"log"
	"net/http"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/entry/types/tree"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
)

type cacheMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	t := s.Trees["cache"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = tools.Title(tools.KineName[m.Lang])
	m.Section = "cache"
	m.Desc = t.Info().Field("description", m.Lang)

	err := m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := s.Recents["cache"].Access(m.Auth.Subscriber)[m.Lang].Limit(10)

	err = s.ExecuteTemplate(w, "cache-main", &cacheMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}
