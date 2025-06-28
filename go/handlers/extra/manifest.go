package extra

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

func Manifest(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Title = meta.SiteName(m.Lang)
		m.Desc = v.Engine.Vars.Lang("site", m.Lang)

		/*
			m.SetHreflang(nil)
			if err != nil {
				log.Println(err)
				return
			}
		*/

		err := v.Engine.ExecuteTemplate(w, "manifest", m)
		if err != nil {
			log.Println(err)
		}
	}
}
