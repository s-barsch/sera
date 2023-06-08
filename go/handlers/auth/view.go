package auth

import (
	"log"
	"net/http"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/meta"
	"strings"
)

type extraHold struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func lastItem(path string) string {
	items := strings.Split(strings.Trim(path, "/"), "/")
	return items[len(items)-1]
}

func SysPage(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	extra := s.Trees["extra"].Access(m.Auth.Subscriber)[m.Lang]

	t, err := extra.SearchTree(lastItem(m.Path), m.Lang)
	if err != nil {
		s.Debug(err)
		http.NotFound(w, r)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	m.Title = t.Title(m.Lang)
	m.Section = "extra"

	err = m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	err = s.ExecuteTemplate(w, t.Slug("en")+"-extra", &extraHold{
		Meta: m,
		Tree: t,
	})
	if err != nil {
		log.Println(err)
	}
}
