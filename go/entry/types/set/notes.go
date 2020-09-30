package set

import (
	"sacer/go/entry/types/text"
)

type Footnotes struct {
	En, De []string
}

func (s *Set) RenderFootnotes() {
	c := 1
	notes := &Footnotes{}
	for _, e := range s.Entries() {
		t, ok := e.(*text.Text)
		if ok {
			t.RenderFootnotes(c)

			notes.De = append(notes.De, t.Notes["de"]...)
			notes.En = append(notes.En, t.Notes["en"]...)

			c += len(t.Notes["de"])
		}
	}

	s.Notes = notes
}

