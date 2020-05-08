package tree

import (
	"stferal/go/entry"
	"stferal/go/entry/types/media/text"
	"stferal/go/entry/types/set"
)


func (ts Trees) Reverse() Trees {
	n := Trees{}
	for i := len(ts) - 1; i >= 0; i-- {
		n = append(n, ts[i])
	}
	return n
}

func (tree *Tree) TraverseTrees() Trees {
	trees := Trees{tree}
	for _, t := range tree.Trees.Reverse() {
		ts := t.TraverseTrees()
		trees = append(trees, ts...)
	}
	return trees
}

func (tree *Tree) TraverseEntries() entry.Entries {
	trees := tree.TraverseTrees()

	entries := entry.Entries{}

	for _, t := range trees {
		entries = append(entries, t.Entries()...)
		//sort.Sort(Desc(h.entries))
	}

	return entries
}

// Trees are traversed in regular order, but entries are reversed.
func (tree *Tree) TraverseEntriesReverse() entry.Entries {
	trees := tree.TraverseTrees()

	es := entry.Entries{}

	for _, t := range trees {
		es = append(es, t.Entries().Reverse()...)
	}

	return es
}

// public functions

func (t *Tree) Public() *Tree {
	nt := t.Copy()
	trees := Trees{}
	for _, tree := range nt.Trees {
		if tree.Info()["private"] == "true" {
			continue
		}
		trees = append(trees, tree.Public())
	}

	nt.entries = makePublic(nt.entries)
	nt.Trees = trees

	return nt
}

func makePublic(es entry.Entries) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if e.Info()["private"] == "true" {
			continue
		}
		s, ok := e.(*set.Set)
		if ok {
			ns := s.Copy()
			ns.SetEntries(makePublic(ns.Entries()))
			e = ns
		}
		l = append(l, e)
	}
	return l
}


// langs

func (t *Tree) Lang(lang string) *Tree {
	nt := t.Copy()
	trees := Trees{}
	for _, tree := range nt.Trees {
		if isNotTranslated(tree, lang) {
			continue
		}
		trees = append(trees, tree.Lang(lang))
	}

	nt.entries = langOnly(nt.entries, lang)
	nt.Trees = trees

	return nt
}

func langOnly(es entry.Entries, lang string) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if isEmptyText(e, lang) {
			continue
		}
		s, ok := e.(*set.Set)
		if ok {
			if isNotTranslated(s, lang) {
				continue
			}
			ns := s.Copy()
			ns.SetEntries(langOnly(ns.Entries(), lang))
			e = ns
		}

		l = append(l, e)
	}
	return l
}

func isNotTranslated(e entry.Entry, lang string) bool {
	_, ok := e.(entry.Collection)
	if !ok {
		return false
	}
	if e.Info().Title(lang) == "" {
		return true
	}
	return lang != "de" && e.Info()["translated"] == "false"
}

func isEmptyText(e entry.Entry, lang string) bool {
	tx, ok := e.(*text.Text)
	if ok {
		if tx.Text(lang) == "" {
			return true
		}
	}
	return false
}
