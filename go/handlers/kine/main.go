package kine

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

type kineMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	t := s.Trees["kine"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = strings.Title(tools.KineName[m.Lang])
	m.Section = "kine"
	m.Desc = s.Vars.Lang("kine-desc", m.Lang)

	err := m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := s.Recents["kine"].Access(m.Auth.Subscriber)[m.Lang].Limit(10)

	err = s.ExecuteTemplate(w, "kine-main", &kineMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}
