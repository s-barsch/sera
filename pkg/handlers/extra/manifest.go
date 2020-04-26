package extra

import (
	"net/http"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

func Manifest(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Desc:    s.Vars.Lang("site", head.Lang(r.Host)),
		Dark:   head.DarkMode(r),
	}
	err := s.ExecuteTemplate(w, "manifest", head)
	if err != nil {
		s.Log.Println(err)
	}
}
