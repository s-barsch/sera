package extra

import (
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
)

func Manifest(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	m.Title = meta.SiteName(m.Lang)
	m.Desc = s.Vars.Lang("site", m.Lang)

	/*
		err := m.Process(nil)
		if err != nil {
			s.Log.Println(err)
			return
		}
	*/

	err := s.ExecuteTemplate(w, "manifest", m)
	if err != nil {
		s.Log.Println(err)
	}
}
