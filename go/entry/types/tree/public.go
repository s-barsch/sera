package tree

import (
	"stferal/go/entry"
	"stferal/go/entry/types/set"
)

func (t *Tree) MakePublic() *Tree {
	trees := Trees{}
	for _, tree := range t.Trees {
		if tree.Info()["private"] == "true" {
			continue
		}
		trees = append(trees, tree.MakePublic())
	}

	t.entries = MakePublicEntries(t.entries)
	t.Trees = trees

	return t
}

func MakePublicEntries(es entry.Entries) entry.Entries {
	l := entry.Entries{}
	for _, e := range es {
		if e.Info()["private"] == "true" {
			continue
		}
		s, ok := e.(*set.Set)
		if ok {
			s.SetEntries(MakePublicEntries(s.Entries()))
			e = s
		}
		l = append(l, e)
	}
	return l
}



