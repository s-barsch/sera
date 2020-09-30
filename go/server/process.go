package server

import (
	"sacer/go/entry"
	"sacer/go/entry/types/set"
	"sacer/go/server/process"
)

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
	kines := s.Recents["kine"].Local(s.Flags.Local)["de"]
	graph := s.Trees["graph"].Local(s.Flags.Local)["de"]

	for _, t := range graph.TraverseTrees() {
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
