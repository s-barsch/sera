package indecs

import (
	"log"
	"net/http"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/types/tree"
	"g.rg-s.com/sacer/go/server"
	"g.rg-s.com/sacer/go/server/meta"
)

type indecsMain struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Recents entry.Entries
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	t := s.Trees["indecs"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = "indecs"
	m.Section = "indecs"

	err := m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	recents := s.Recents["indecs"].Access(m.Auth.Subscriber)[m.Lang]

	err = s.ExecuteTemplate(w, "indecs-main", &indecsMain{
		Meta:    m,
		Tree:    t,
		Recents: recents.Offset(0, 100),
	})
	if err != nil {
		log.Println(err)
	}
}
