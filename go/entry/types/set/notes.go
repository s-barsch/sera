package set

import (
	"sacer/go/entry"
	"sacer/go/entry/types/text"
)

func NumberFootnotes(entries entry.Entries) text.Notes {
	c := 1
	notes := map[string][]string{}
	for _, e := range entries {
		t, ok := e.(*text.Text)
		if ok {
			t.Script.NumberFootnotes(c)

			notes["de"] = append(notes["de"], t.Script.Notes["de"]...)
			notes["en"] = append(notes["en"], t.Script.Notes["en"]...)

			c += len(t.Script.Notes["de"])
		}
	}

	return notes
}

