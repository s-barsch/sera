package index

import (
	"log"
	"net/http"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type indexSerial struct {
	Head *head.Head
	Els  entry.Els
}

func Serial(s *server.Server, w http.ResponseWriter, r *http.Request) {
	h := &head.Head{
		Title:   "Serial - Index",
		Section: "index",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      nil,
		Desc:    s.Vars.Lang("serial", head.Lang(r.Host)),
		Night:   head.NightMode(r),
		Large:   head.TypeMode(r),
		NoLog:   head.LogMode(r),
	}
	err := h.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	//head.Description = srv.varsLang("serial", lang)

	h.Langs = []*head.Link{
		&head.Link{
			Name: "de",
			Href: h.AbsoluteURL("/index/serial/", "de"),
		},
		&head.Link{
			Name: "en",
			Href: h.AbsoluteURL("/index/serial/", "en"),
		},
	}

	recents := s.Recents["index"].NoEmpty(h.Lang)

	if !s.Flags.Local {
		recents = recents.ExcludePrivate()
	}

	err = s.ExecuteTemplate(w, "index-serial", &indexSerial{
		Head: h,
		Els: recents,
	})
	if err != nil {
		log.Println(err)
	}
}
