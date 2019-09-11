package index 

import (
	"log"
	"net/http"
	"stferal/pkg/el"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

type holdPage struct {
	Head *head.Head
	Hold *el.Hold
}

func Hold(s *server.Server, w http.ResponseWriter, r *http.Request, hold *el.Hold) {
	if perma := hold.Permalink(head.Lang(r.Host)); r.URL.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}
	head := &head.Head{
		Title:   holdTitle(hold, head.Lang(r.Host)),
		Section: "index",
		Path:    r.URL.Path,
		Host:    r.Host,
		El:      hold,
	}
	err := head.Make()
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, "index-hold", &holdPage{
		Head: head,
		Hold: hold,
	})
	if err != nil {
		log.Println(err)
	}
}

func holdTitle(hold *el.Hold, lang string) string {
	title := hold.Info.Title(lang)

	c := hold.Chain(lang)
	if len(c) > 2 {
		title += " - " + c[1].Title
	}

	return title
}
