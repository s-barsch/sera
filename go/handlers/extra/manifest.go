package extra

import (
	"net/http"
	"stferal/go/head"
	"stferal/go/server"
)

func Manifest(s *server.Server, w http.ResponseWriter, r *http.Request) {
	head := &head.Head{
		Desc: s.Vars.Lang("site", head.Lang(r.Host)),
		Options: head.GetOptions(r),
	}
	err := s.ExecuteTemplate(w, "manifest", head)
	if err != nil {
		s.Log.Println(err)
	}
}
