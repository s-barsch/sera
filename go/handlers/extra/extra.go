package extra

import (
	"log"
	"net/http"
	"strings"

	"g.rg-s.com/sera/go/entry/types/tree"
	"g.rg-s.com/sera/go/server/meta"
	"g.rg-s.com/sera/go/viewer"
)

type extraHold struct {
	Meta *meta.Meta
	Tree *tree.Tree
}

func lastItem(path string) string {
	items := strings.Split(strings.Trim(path, "/"), "/")
	return items[len(items)-1]
}

func Extra(v *viewer.Viewer, m *meta.Meta) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		extra := v.Store.Trees["extra"].Access(m.Auth.Subscriber)[m.Lang]
		t, err := extra.SearchTree(lastItem(m.Path), m.Lang)
		if err != nil {
			v.Logger.Info(err)
			http.NotFound(w, r)
			return
		}

		if perma := t.Perma(m.Lang); m.Path != perma {
			http.Redirect(w, r, perma, http.StatusMovedPermanently)
			return
		}

		m.Title = t.Title(m.Lang)
		m.SetSection("extra")
		m.SetHreflang(t)

		err = v.Engine.ExecuteTemplate(w, "extra-page", &extraHold{
			Meta: m,
			Tree: t,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
