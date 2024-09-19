package set

import (
	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools/script"
	"g.rg-s.com/sera/go/entry/types/text"
)

func NumberFootnotes(entries entry.Entries) script.Footnotes {
	c := 1
	notes := map[string][]string{}
	for _, e := range entries {
		t, ok := e.(*text.Text)
		if ok {
			t.Script.NumberFootnotes(c)

			notes["de"] = append(notes["de"], t.Script.Footnotes["de"]...)
			notes["en"] = append(notes["en"], t.Script.Footnotes["en"]...)

			c += len(t.Script.Footnotes["de"])
		}
	}

	return notes
}
