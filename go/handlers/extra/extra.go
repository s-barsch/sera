package extra

import (
	"log"
	"net/http"
	"strings"

	"g.rg-s.com/sera/go/entry/types/tree"
	s "g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/meta"
)

type extraHold struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func lastItem(path string) string {
	items := strings.Split(strings.Trim(path, "/"), "/")
	return items[len(items)-1]
}

func Extra(w http.ResponseWriter, r *http.Request, m *meta.Meta) {

	extra := s.Store.Trees["extra"].Access(m.Auth.Subscriber)[m.Lang]
	t, err := extra.SearchTree(lastItem(m.Path), m.Lang)
	if err != nil {
		s.Store.Debug(err)
		http.NotFound(w, r)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, http.StatusMovedPermanently)
		return
	}

	m.Title = t.Title(m.Lang)
	m.Section = "extra"

	err = m.Process(t)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Store.ExecuteTemplate(w, "extra-page", &extraHold{
		Meta: m,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}
