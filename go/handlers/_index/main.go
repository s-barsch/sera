package index

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type indexMain struct {
	Meta   *meta.Meta
	Phanes *tree.Tree
}

func Main(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	komposita := s.Trees["komposita"].Access(m.Auth.Subscriber)[m.Lang]

	m.Title = "Index"

	m.SetSection("index")
	m.SetHreflang(nil)

	err = s.ExecuteTemplate(w, "index-main", &indexMain{
		Meta:   m,
		Phanes: komposita,
	})
	if err != nil {
		log.Println(err)
	}
}
