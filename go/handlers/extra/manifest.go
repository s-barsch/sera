package extra

import (
	"log"
	"net/http"

	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func Manifest(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	m.Title = meta.SiteName(m.Lang)
	m.Desc = s.Srv.Vars.Lang("site", m.Lang)

	/*
		m.SetHreflang(nil)
		if err != nil {
			log.Println(err)
			return
		}
	*/

	err := s.Srv.ExecuteTemplate(w, "manifest", m)
	if err != nil {
		log.Println(err)
	}
}
