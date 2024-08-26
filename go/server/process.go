package server

import (
	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/entry/types/set"
)

func (s *Server) makeLinks() {
	reels := s.Recents["reels"].Access(true)["de"]

	for _, access := range []bool{false, true} {
		for lang := range tools.Langs {
			for _, e := range s.Recents["graph"].Access(access)[lang] {
				s, ok := e.(*set.Set)
				if ok {
					es := findMatchingKines(reels, s)
					if es != nil {
						s.Kine = es
					}
				}
			}
		}
	}
}

func findMatchingKines(reels entry.Entries, s *set.Set) entry.Entries {
	matches := entry.Entries{}
	for _, e := range reels {
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
