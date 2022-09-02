package index

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/meta"
)

type indexMain struct {
	Meta   *meta.Meta
	Phanes *tree.Tree
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	komposita := s.Trees["komposita"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = "Index"
	m.Section = "index"

	err := m.Process(nil)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "index-main", &indexMain{
		Meta:   m,
		Phanes: komposita,
	})
	if err != nil {
		log.Println(err)
	}
}
