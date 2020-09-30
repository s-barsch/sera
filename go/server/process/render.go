package process

import (
	"sacer/go/entry"
	/*
	"sacer/go/entry/info"
	"sacer/go/entry/types/text"
	"sacer/go/entry/types/set"
	"sacer/go/entry/types/tree"
	*/
	"sacer/go/entry/tools/hyph"
)


type HyphPatterns map[string]*hyph.Patterns

func (patterns HyphPatterns) HyphEntries(entries entry.Entries) {
	/*
	for _, e := range entries {
		patterns.HyphInfo(e)
		s, ok := e.(*set.Set)
		if ok {
			patterns.HyphEntries(s.Entries())
			continue
		}
		t, ok := e.(*tree.Tree)
		if ok {
			patterns.HyphEntries(t.Entries())
			continue
		}
		tx, ok := e.(*text.Text)
		if ok {
			patterns.HyphTextEntry(tx)
		}
	}
*/
}


/*
func (patterns HyphPatterns) HyphInfo(e entry.Entry) {
	inf := e.Info()
	for _, key := range []string{"title", "transcript"} {
		inf = patterns.HyphInfoField(inf, key)
	}
	e.SetInfo(inf)
}

func (patterns HyphPatterns) HyphInfoField(inf info.Info, key string) info.Info {
	setKey := key
	if key == "title" {
		setKey = "title-hyph"
	}
	for _, l := range langs {
		if l == "en" {
			setKey += "-en"
		}
		inf[setKey] = patterns[l].HyphenateText(inf.Field(key, l))
	}
	return inf
}

func (patterns HyphPatterns) HyphTextEntry(tx *text.Text) {
	for _, l := range langs {
		tx.TextLangs[l] = patterns[l].HyphenateText(tx.TextLangs[l])
	}
}

*/
