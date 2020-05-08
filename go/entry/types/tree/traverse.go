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
	trees := Trees{}
	for _, tree := range t.Trees {
		if tree.Info()["private"] == "true" {
			continue
		}
		trees = append(trees, tree.Public())
	}

	t.entries = makePublic(t.entries)
	t.Trees = trees

	return t
}

func makePublic(es entry.Entries) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if e.Info()["private"] == "true" {
			continue
		}
		s, ok := e.(*set.Set)
		if ok {
			s.SetEntries(makePublic(s.Entries()))
			e = s
		}
		l = append(l, e)
	}
	return l
}


// langs

func (t *Tree) Lang(lang string) *Tree {
	trees := Trees{}
	for _, tree := range t.Trees {
		if lang == "de" && tree.Info()["translated"] == "false" {
			continue
		}
		trees = append(trees, tree.Public())
	}

	t.entries = langOnly(t.entries, lang)
	t.Trees = trees

	return t
}

func langOnly(es entry.Entries, lang string) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if e.Info()["private"] == "true" {
			continue
		}
		tx, ok := e.(*text.Text)
		if ok {
			if tx.Text(lang) == "" {
				continue
			}
		}
		l = append(l, e)
	}
	return l
}


