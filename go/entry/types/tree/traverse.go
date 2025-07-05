package tree

import (
	"cmp"
	"slices"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/types/set"
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

func (tree *Tree) Flatten() entry.Entries {
	trees := tree.TraverseTrees()

	entries := []entry.Entry{}

	for _, t := range trees {
		entries = append(entries, t.Entries()...)
	}

	slices.SortFunc(entries, func(a, b entry.Entry) int {
		return cmp.Compare(a.Id(), b.Id())
	})

	return entries
}

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
