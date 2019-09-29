package index 

import (
	"log"
	"net/http"
	"stferal/pkg/el"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type indexSerial struct {
	Head *head.Head
	Els el.Els
}

func Serial(s *server.Server, w http.ResponseWriter, r *http.Request) {
	h:= &head.Head{
		Title:   "Serial - Index",
		Section: "index",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      nil,
		Desc:    s.Vars.Lang("serial", head.Lang(r.Host)),
		Dark:    head.DarkMode(r),
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

	err = s.ExecuteTemplate(w, "index-serial", &indexSerial{
		Head: h,
		Els: s.Recents["index"],
	})
	if err != nil {
		log.Println(err)
	}
}
