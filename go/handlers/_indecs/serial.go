package indecs

import (
	"log"
	"net/http"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/server"
	"g.rg-s.com/sacer/go/server/meta"
)

type indecsSerial struct {
	meta    *meta.Meta
	Entries entry.Entries
}

func Serial(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	m.Title = "Serial - Index"
	m.Section = "indecs"

	//Desc:    s.Vars.Lang("serial", meta.Lang(r.Host)),

	err := m.Process(nil)
	if err != nil {
		s.Log.Println(err)
		return
	}

	m.Langs = []*meta.Link{
		{
			Name: "de",
			Href: m.AbsoluteURL("/indecs/serial", "de"),
		},
		{
			Name: "en",
			Href: m.AbsoluteURL("/indecs/serial", "en"),
		},
	}

	recents := s.Recents["indecs"].Access(m.Auth.Subscriber)[m.Lang]

	err = s.ExecuteTemplate(w, "indecs-serial", &indecsSerial{
		meta:    m,
		Entries: recents,
	})
	if err != nil {
		log.Println(err)
	}
}
