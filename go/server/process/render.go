package process

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/parts/info"
	"stferal/go/entry/types/media/text"
	"stferal/go/entry/types/set"
	"stferal/go/entry/types/tree"
	"stferal/go/entry/helper/hyph"
)

type HyphPatterns map[string]*hyph.Patterns

func LoadHyphPatterns(root string) (HyphPatterns, error) {
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

var langs = []string{
	"de",
	"en",
}

func RenderTexts(root string, entries entry.Entries) error {
	h, err := LoadHyphPatterns(root)
	if err != nil {
		return err
	}
	h.HyphenateEntries(entries)
	return nil
}


func (h HyphPatterns) HyphenateEntries(entries entry.Entries) {
	for _, e := range entries {
		h.HyphenateFields(e)
		s, ok := e.(*set.Set)
		if ok {
			h.HyphenateEntries(s.Entries())
			continue
		}
		t, ok := e.(*tree.Tree)
		if ok {
			h.HyphenateEntries(t.Entries())
			continue
		}
		tx, ok := e.(*text.Text)
		if !ok {
			continue
		}
		h.HyphenateTextEntry(tx)
	}
}

/*
func (h HyphPatterns) HyphenateTitles(es entry.Entries) {
	for _, e := range es {
		h.HyphenateTitle(e)
	}
}
*/
func (h HyphPatterns) HyphenateFields(e entry.Entry) {
	inf := e.Info()
	for _, key := range []string{"title", "transcript"} {
		inf = h.HyphenateField(inf, key)
	}
	e.SetInfo(inf)
}

func (h HyphPatterns) HyphenateField(inf info.Info, key string) info.Info {
	setKey := key
	if key == "title" {
		setKey = "title-hyph"
	}
	for _, l := range langs {
		if l == "en" {
			setKey += "-en"
		}
		inf[setKey] = h[l].HyphenateText(inf.Field(key, l))
	}
	return inf
}

func (h HyphPatterns) HyphenateTextEntry(tx *text.Text) {
	for _, l := range langs {
		tx.TextLangs[l] = h[l].HyphenateText(tx.TextLangs[l])
	}
}

