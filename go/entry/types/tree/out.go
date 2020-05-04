package tree

import (
	"fmt"
	"stferal/go/entry"
)

// For this to work, title fields have to be set for all sections. Otherwise,
// the root node will return a short hash.
func (t *Tree) Section() string {
	return t.Chain()[0].Slug("en")
}

// /index/welt/wuestenleben-36c35dcb
func (t *Tree) Perma(lang string) string {
	return fmt.Sprintf("%v-%v", t.Path(lang), t.Hash())
}

// /index/welt/wuestenleben
func (t *Tree) Path(lang string) string {
	path := ""
	for _, tree := range t.Chain() {
		path += "/" + tree.Slug(lang)
	}
	return path
}

func (t *Tree) Chain() Trees {
	parent := typeCheck(t.Parent())
	if parent == nil {
		return Trees{t}
	}

	return append(parent.Chain(), t)
}

func typeCheck(parentEntry entry.Entry) *Tree {
	parent, ok := parentEntry.(*Tree)
	if !ok {
		return nil
	}

	return parent
}

func (t *Tree) Level() int {
	parent := typeCheck(t.Parent())
	if parent == nil {
		return 0
	}
	return 1 + parent.Level()
}


