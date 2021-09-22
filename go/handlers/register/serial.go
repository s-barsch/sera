package register

import (
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/server"
	"sacer/go/server/auth"
	"sacer/go/server/head"
)

type registerSerial struct {
	Head    *head.Head
	Entries entry.Entries
}

func Serial(s *server.Server, w http.ResponseWriter, r *http.Request, a *auth.Auth) {
	lang := head.Lang(r.Host)
	h := &head.Head{
		Title:   "Serial - Index",
		Section: "register",
		Path:    r.URL.Path,
		Host:    r.Host,
		Entry:   nil,
		//Desc:    s.Vars.Lang("serial", head.Lang(r.Host)),
		Options: head.GetOptions(r),
	}
	err := h.Process()
	if err != nil {
		s.Log.Println(err)
		return
	}

	//head.Description = srv.varsLang("serial", lang)

	h.Langs = []*head.Link{
		&head.Link{
			Name: "de",
			Href: h.AbsoluteURL("/register/serial", "de"),
		},
		&head.Link{
			Name: "en",
			Href: h.AbsoluteURL("/register/serial", "en"),
		},
	}

	recents := s.Recents["register"].Access(a.Subscriber)[lang]

	err = s.ExecuteTemplate(w, "register-serial", &registerSerial{
		Head:    h,
		Entries: recents,
	})
	if err != nil {
		log.Println(err)
	}
}
