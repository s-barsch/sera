package server

import (
	"sacer/go/entry"
	"sacer/go/entry/types/set"
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

	err = s.processAllTexts()
	if err != nil {
		return err
	}

	s.makeLinks()

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
	// Hyphenation will be done for both languages. We just iterate through
	// the German tree here, but will process all fiels within the entries.
	lang := "de"
	for _, section := range sections {
		for _, e := range s.Trees[section].Private[lang].TraverseTrees() {
			patterns.HyphInfo(e)
		}
		patterns.HyphEntries(s.Recents[section].Private[lang])
	}
	return nil
}

func (s *Server) makeLinks() {
	kines := s.Recents["kine"].Private["de"]

	for _, t := range s.Trees["graph"].Private["de"].TraverseTrees() {
		for _, e := range t.Entries() {
			s, ok := e.(*set.Set)
			if ok {
				es := findMatchingKines(kines, s)
				if es != nil {
					s.Kine = es
				}
			}
		}
	}
}

func findMatchingKines(kines entry.Entries, s *set.Set) entry.Entries {
	matches := entry.Entries{}
	for _, e := range kines {
		if e.Date().Format("060102") == s.Date().Format("060102") {
			matches = append(matches, e)
		}
	}
	if len(matches) > 0 {
		return matches
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
