package process

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/types/media/text"
	"stferal/go/entry/types/set"
	"stferal/go/entry/helper/hyph"
)

var langs = []string{
	"de",
	"en",
}

func RenderTexts(root string, entries entry.Entries) error {
	hs, err := LoadHyphPatterns(root)
	if err != nil {
		return err
	}
	ProcessEntries(hs, entries)
	return nil
}

func ProcessEntries(hs map[string]*hyph.Patterns, entries entry.Entries) {
	for _, e := range entries {
		s, ok := e.(*set.Set)
		if ok {
			ProcessEntries(hs, s.Entries())
			continue
		}
		tx, ok := e.(*text.Text)
		if !ok {
			continue
		}
		ProcessText(hs, tx)
	}
}

func LoadHyphPatterns(root string) (map[string]*hyph.Patterns, error) {
	hs := map[string]*hyph.Patterns{}
	for _, l := range langs {
		p, err := hyph.LoadPattern(fmt.Sprintf("%v/go/entry/helper/hyph/hyph-%v.dic", root, l))
		if err != nil {
			return nil, err
		}
		hs[l] = p
	}
	return hs, nil
}

func ProcessText(hs map[string]*hyph.Patterns, tx *text.Text) {
	for _, l := range langs {
		tx.TextLangs[l] = hs[l].HyphenateText(tx.TextLangs[l])
	}
}

