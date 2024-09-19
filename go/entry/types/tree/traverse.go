package tree

import (
	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/image"
	"g.rg-s.com/sera/go/entry/types/set"
	"g.rg-s.com/sera/go/entry/types/text"
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
		ntree := tree.Public()
		ntree.parent = nt
		trees = append(trees, ntree)
	}

	nt.entries = makePublic(nt.entries, nt)
	nt.Trees = trees

	return nt
}

func makePublic(es entry.Entries, newParent entry.Entry) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if e.Info()["private"] == "true" {
			continue
		}
		e.SetParent(newParent)
		if s, ok := e.(*set.Set); ok {
			ns := s.Copy()
			ns.SetEntries(makePublic(ns.Entries(), ns))
			ns.SetNotes(makePublic(ns.Notes, ns))
			e = ns
		}
		l = append(l, e)
	}
	return l
}

// blur

func (t *Tree) Blur() *Tree {
	nt := t.Copy()
	trees := Trees{}
	if nt.Info().Wall() {
		nt.Info()["blur"] = "true"
	}
	for _, child := range nt.Trees {
		if nt.Info().Wall() {
			child.Info()["wall"] = "true"
		}
		nchild := child.Blur()
		nchild.parent = nt
		trees = append(trees, nchild)
	}

	nt.entries = makeBlur(nt.entries, nt)
	nt.Trees = trees

	return nt
}

func makeBlur(es entry.Entries, newParent entry.Entry) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if newParent.Info().Wall() || e.Info().Wall() {
			e.SetParent(newParent)
			if s, ok := e.(*set.Set); ok {
				ns := s.Copy()
				e.Info()["wall"] = "true"
				ns.SetEntries(makeBlur(ns.Entries(), ns))

				if ns.Cover != nil {
					ns.Cover = ns.Cover.Blur()
				}

				e = ns
			}
			if t, ok := e.(*text.Text); ok {
				e = t.Blur()
			}
			if i, ok := e.(*image.Image); ok {
				e = i.Blur()
			}
			e.Info()["blur"] = "true"
		}
		l = append(l, e)
	}
	return l
}

// langs

func (t *Tree) Translated(lang string) *Tree {
	nt := t.Copy()
	trees := Trees{}
	for _, tree := range nt.Trees {
		if isNotTranslated(tree, lang) {
			continue
		}
		ntree := tree.Translated(lang)
		ntree.parent = nt
		trees = append(trees, ntree)
	}

	nt.entries = langOnly(nt.entries, lang, nt)
	nt.Trees = trees

	return nt
}

func langOnly(es entry.Entries, lang string, newParent entry.Entry) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if isEmptyText(e, lang) {
			// Formerly: continue.
			// Empty texts are now kept and displayed via html template.
		}
		e.SetParent(newParent)
		s, ok := e.(*set.Set)
		if ok {
			if isNotTranslated(s, lang) {
				continue
			}
			ns := s.Copy()
			ns.SetEntries(langOnly(ns.Entries(), lang, ns))
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
	/*
		if e.Info().Title(lang) == "" {
			return true
		}
	*/
	return lang != "de" && e.Info()["translated"] == "false"
}

func isEmptyText(e entry.Entry, lang string) bool {
	tx, ok := e.(*text.Text)
	if ok {
		if tx.Script.Langs[lang] == "" {
			return true
		}
	}
	return false
}
