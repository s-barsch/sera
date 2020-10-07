package extra

import (
	"net/http"
	"sacer/go/head"
	"sacer/go/server"
	"sacer/go/server/auth"
)

func Manifest(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	lang := head.Lang(r.Host)
	head := &head.Head{
		Title:   head.SiteName(lang),
		Desc:    s.Vars.Lang("site", lang),
		Options: head.GetOptions(r),
		Host:    r.Host,
	}
	err := head.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "manifest", head)
	if err != nil {
		s.Log.Println(err)
	}
}
