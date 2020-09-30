package set

import (
	"sacer/go/entry/types/text"
)

func (s *Set) RenderFootnotes() {
	c := 1
	notes := map[string][]string{}
	for _, e := range s.Entries() {
		t, ok := e.(*text.Text)
		if ok {
			t.RenderFootnotes(c)

			notes["de"] = append(notes["de"], t.Notes["de"]...)
			notes["en"] = append(notes["en"], t.Notes["en"]...)

			c += len(t.Notes["de"])
		}
	}

	s.Notes = notes
}

