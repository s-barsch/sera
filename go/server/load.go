package server

import (
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"stferal/go/server/tmpl"
	"time"
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

	trees := map[string]*tree.Tree{}
	recents := map[string]entry.Entries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}

		sPrivate := section+"-private"
		sPublic := section

		trees[sPrivate] = t
		trees[sPublic] = t.MakePublic()

		recents[sPrivate] = serialize(trees[sPrivate])
		recents[sPublic] = serialize(trees[sPublic])
	}

	s.Trees = trees
	s.Recents = recents

	return nil
}

func serialize(t *tree.Tree) entry.Entries {
	if t.Section() == "graph" {
		return t.TraverseEntriesReverse()
	}
	return t.TraverseEntries().Exclude().Desc()
}


