package server

import (
	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/types/set"
)

func (s *Server) makeLinks() {
	cache := s.Recents["cache"].Access(true)["de"]

	for _, access := range []bool{false, true} {
		for lang := range tools.Langs {
			for _, e := range s.Recents["graph"].Access(access)[lang] {
				s, ok := e.(*set.Set)
				if ok {
					es := findMatchingKines(cache, s)
					if es != nil {
						s.Kine = es
					}
				}
			}
		}
	}
}

func findMatchingKines(cache entry.Entries, s *set.Set) entry.Entries {
	matches := entry.Entries{}
	for _, e := range cache {
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
