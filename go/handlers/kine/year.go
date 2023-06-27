package kine

import (
	"fmt"
	"log"
	"net/http"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sacer/go/server"
	"sacer/go/server/meta"
	"sacer/go/server/paths"
	"strings"
	"time"
)

type kineYear struct {
	Meta    *meta.Meta
	Tree    *tree.Tree
	Entries entry.Entries
}

func Year(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta, p *paths.Path) {
	kine := s.Trees["kine"].Access(m.Auth.Subscriber)[m.Lang]

	id, err := getYearId(p.Slug)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	t, err := kine.LookupTree(id)
	if err != nil {
		http.NotFound(w, r)
		s.Log.Println(err)
		return
	}

	if perma := t.Perma(m.Lang); m.Path != perma {
		http.Redirect(w, r, perma, 301)
		return
	}

	m.Title = strings.Title(fmt.Sprintf("%v - %v", t.Date().Format("2006"), tools.KineName[m.Lang]))
	m.Section = "kine"
	// TODO:
	//m.Desc = s.Vars.Lang("kine-desc", m.Lang)

	err = m.Process(t)
	if err != nil {
		s.Log.Println(err)
		return
	}

	entries := t.TraverseEntriesReverse()

	err = s.ExecuteTemplate(w, "kine-year", &kineMain{
		Meta:    m,
		Tree:    t,
		Entries: entries,
	})
	if err != nil {
		log.Println(err)
	}
}

func getYearId(year string) (int64, error) {
	t, err := time.Parse("2006", year)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
