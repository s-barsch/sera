package server

import (
	"sacer/go/entry"
	"sacer/go/entry/types/set"
)

func (s *Server) makeLinks() {
	kines := s.Recents["kine"]["de"]
	graph := s.Trees["graph"]["de"]

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
		// TODO: start at 20-08
		if e.Date().Format("060102") == s.Date().Format("060102") {
			matches = append(matches, e)
		}
	}
	if len(matches) > 0 {
		return matches
	}
	return nil
}
