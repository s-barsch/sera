package index

import (
	"log"
	"net/http"
	"stferal/go/entry"
	"stferal/go/head"
	"stferal/go/server"
)

type indexSerial struct {
	Head *head.Head
	Entries  entry.Entries
}

func Serial(s *server.Server, w http.ResponseWriter, r *http.Request) {
	lang := head.Lang(r.Host)
	h := &head.Head{
		Title:   "Serial - Index",
		Section: "index",
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
			Href: h.AbsoluteURL("/index/serial", "de"),
		},
		&head.Link{
			Name: "en",
			Href: h.AbsoluteURL("/index/serial", "en"),
		},
	}

	recents := s.Recents["index"].Local(s.Flags.Local)[lang]

	err = s.ExecuteTemplate(w, "index-serial", &indexSerial{
		Head:    h,
		Entries: recents,
	})
	if err != nil {
		log.Println(err)
	}
}
