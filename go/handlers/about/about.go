package about

import (
	"log"
	"net/http"

	"g.rg-s.com/sacer/go/entry/types/tree"
	"g.rg-s.com/sacer/go/server"
	"g.rg-s.com/sacer/go/server/meta"
)

type aboutTree struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func ServeAbout(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, t *tree.Tree) {
	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	m.Title = t.Title(m.Lang)
	m.Section = "about"

	err := m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, aboutTemplate(t.Level()), &aboutTree{
		Meta: m,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}

func aboutTemplate(level int) string {
	if level == 0 {
		return "about-main"
	}
	return "about-page"
}
