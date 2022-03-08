package server

import (
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/set"
)

func (s *Server) makeLinks() {
	kinos := s.Recents["kino"].Access(true)["de"]

	for _, access := range []bool{false, true} {
		for lang, _ := range tools.Langs {
			for _, e := range s.Recents["graph"].Access(access)[lang] {
				s, ok := e.(*set.Set)
				if ok {
					es := findMatchingKinos(kinos, s)
					if es != nil {
						s.Kino = es
					}
				}
			}
		}
	}
}

func findMatchingKinos(kinos entry.Entries, s *set.Set) entry.Entries {
	matches := entry.Entries{}
	for _, e := range kinos {
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
