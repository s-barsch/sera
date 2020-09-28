package server

import (
	"sacer/go/server/process"
	"sacer/go/server/tmpl"
	"time"
)

var sections = []string{
	"index",
	"graph",
	"about",
	"extra",
	"kine",
}

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

	/*
	err = s.processAllTexts()
	if err != nil {
		return err
	}
	*/

	tEnd := time.Now()
	tDif := tEnd.Sub(tStart)

	if s.Flags.Debug {
		s.Log.Printf("Load: %v.\n", tDif)
	}

	return nil
}

func (s *Server) processAllTexts() error {
	patterns, err := process.LoadHyphPatterns(s.Paths.Root)
	if err != nil {
		return err
	}
	// TODO: only German?
	lang := "de"
	for _, section := range sections {
		for _, e := range s.Trees[section].Private[lang].TraverseTrees() {
			patterns.HyphInfo(e)
		}
		patterns.HyphEntries(s.Recents[section].Private[lang])
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
