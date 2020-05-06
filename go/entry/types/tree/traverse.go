package tree

import (
	"stferal/go/entry"
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

/*
func (hold *Hold) TraverseElsReverse() list.Els {
	trees := hold.TraverseTrees()

	els := list.Els{}

	for _, h := range trees {
		els = append(els, h.Els.Reverse()...)
	}

	return els
}
*/

/*
func newEls(els list.Els) list.Els {
	nels := list.Els{}
	for _, e := range els {
		nels = append(nels, e)
	}
	return nels
}
*/

/*
func (hold *Hold) traverseEls(stack []*Hold) list.Els {
	els := list.Els{}
	for _, e := range hold.Els {
		els = append(els, e)
	}
	for i, h := range hold.Trees.Reverse() {
		if i == 0 {
			return append(els, h.traverseEls(append(stack, hold.Trees[1:]...))...)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return append(els, h.traverseEls(stack[1:])...)
		}
	}
	return els
}
*/
