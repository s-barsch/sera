package indecs

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type indecsSerial struct {
	meta    *meta.Meta
	Entries entry.Entries
}

func Serial(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	m.Title = "Serial - Index"

	//Desc:    s.Vars.Lang("serial", meta.Lang(r.Host)),

	m.SetSection("indecs")
	m.SetHreflang(nil)

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

	recents := s.Store.Recents["indecs"].Access(m.Auth.Subscriber)[m.Lang]

	err = s.ExecuteTemplate(w, "indecs-serial", &indecsSerial{
		meta:    m,
		Entries: recents,
	})
	if err != nil {
		log.Println(err)
	}
}
