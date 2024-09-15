package set

import (
	"g.rg-s.com/sferal/go/entry"
)

func (s *Set) SetEntries(es entry.Entries) {
	s.entries = es
}

func (s *Set) SetNotes(es entry.Entries) {
	s.Notes = es
}
