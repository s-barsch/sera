package extra

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

func Manifest(w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	m.Title = meta.SiteName(m.Lang)
	m.Desc = server.Store.Vars.Lang("site", m.Lang)

	/*
		err := m.Process(nil)
		if err != nil {
			log.Println(err)
			return
		}
	*/

	err := server.Store.ExecuteTemplate(w, "manifest", m)
	if err != nil {
		log.Println(err)
	}
}
