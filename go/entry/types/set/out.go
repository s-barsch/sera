package set

import (
	"stferal/go/entry"
)

func (s *Set) SetEntries(es entry.Entries) {
	s.entries = es
}
