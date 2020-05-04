package entry

import (
	"fmt"
)

func (s *Set) Permalink(lang string) string {
	switch name := s.File.Base(); name {
	case "about", "legal", "ueber", "impressum":
		return fmt.Sprintf("/%v/", s.Info.Slug(lang))
	}
	return fmt.Sprintf("%v/%v-%v", s.File.Hold.Path(lang), s.Info.Slug(lang), s.Acronym())
}

func (sets Sets) Limit(n int) Sets {
	if len(sets) <= n {
		return sets
	}
	return sets[:n]
}
