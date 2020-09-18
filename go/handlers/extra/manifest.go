package extra

import (
	"net/http"
	"stferal/go/head"
	"stferal/go/server"
)

func Manifest(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)
	head := &head.Head{
		Title:   head.SiteName(lang),
		Desc:    s.Vars.Lang("site", lang),
		Options: head.GetOptions(r),
	}
	err := s.ExecuteTemplate(w, "manifest", head)
	if err != nil {
		s.Log.Println(err)
	}
}
