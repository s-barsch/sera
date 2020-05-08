package server

import (
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"stferal/go/server/tmpl"
	"time"
	"fmt"
)

func (s *Server) Load() error {
	tStart := time.Now()

	err := s.LoadTemplates()
	if err != nil {
		return err

	}

	err = s.LoadTrees()
	if err != nil {
		return err
	}

	tEnd := time.Now()
	tDif := tEnd.Sub(tStart)

	if s.Flags.Debug {
		s.Log.Printf("Load: %v.\n", tDif)
	}

	return nil
}

// templates

func (s *Server) LoadTemplates() error {
	vars, err := tmpl.LoadVars(s.Paths.Root)
	if err != nil {
		return err
	}

	ts, err := tmpl.LoadTemplates(s.Paths.Root, s.Funcs())
	if err != nil {
		return err
	}

	s.Templates = ts
	s.Vars = vars

	return nil
}

// trees

func (s *Server) LoadTrees() error {
	sections := []string{
		"index",
		"graph",
		"about",
		"extra",
	}

	trees := map[string]*SectionTree{}
	recents := map[string]*SectionEntries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}

		trees[section] = &SectionTree{
			Private: makeLangs(t),
			Public:  makeLangs(t.Public()),
		}

		recents[section] = &SectionEntries{
			Private: serializeLangs(trees[section].Private),
			Public:  serializeLangs(trees[section].Public),
		}

		fmt.Printf("%v: lang de %v\n", section, len(recents[section].Public["de"]))
		fmt.Printf("%v: lang en %v\n", section, len(recents[section].Public["en"]))
	}

	s.Trees = trees
	s.Recents = recents

	return nil
}

func makeLangs(t *tree.Tree) map[string]*tree.Tree {
	return map[string]*tree.Tree{
		"de": t,
		"en": t.Lang("en"),
	}
}

func serializeLangs(langMap map[string]*tree.Tree) map[string]entry.Entries {
	return map[string]entry.Entries{
		"de": serialize(langMap["de"]),
		"en": serialize(langMap["en"]),
	}
}

func serialize(t *tree.Tree) entry.Entries {
	if t.Section() == "graph" {
		return t.TraverseEntriesReverse()
	}
	return t.TraverseEntries().Exclude().Desc()
}


/*
func (els Els) NoEmpty(lang string) Els {
	l := Els{}
	for _, e := range els {
		if Type(e) == "text" {
			if e.(*Text).Text[lang] == "" {
				continue
			}
		}
		if Type(e) == "set" && lang != "de" {
			if e.(*Set).Info["translated"] == "false" {
				continue
			}
		}
		l = append(l, e)
	}
	return l
}
*/
