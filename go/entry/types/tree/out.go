package tree

import (
	"fmt"
	"stferal/go/entry"
)

func (t *Tree) Label(lang string) string {
	if label := t.info.Field("label", lang); label != "" {
		return label
	}
	return t.Title(lang)
}

// For this to work, title fields have to be set for all sections. Otherwise,
// the root node will return a short hash.
func (t *Tree) Section() string {
	return t.Chain()[0].Slug("en")
}

func (t *Tree) Perma(lang string) string {
	switch t.Section() {
	case "graph":
		return graphPerma(t, lang)
	case "index":
		return indexPerma(t, lang)
	}
	return fmt.Sprintf("/%v", t.Slug(lang))
}

func extraPerma(t *Tree, lang string) string {
	return fmt.Sprintf("/%v", t.Slug(lang))
}

func graphPerma(t *Tree, lang string) string {
	switch l := t.Level(); {
	case l == 0:
		return "/graph"
	case l == 2:
		return monthAnchor(t.Path(lang))
	case l < 3:
		return t.Path(lang)
	}
	return fmt.Sprintf("/graph-permalink-error-%v", t.Slug(lang))
}

func monthAnchor(path string) string {
	if len(path) > 3 {
		month := len(path) - 3
		return path[:month] + "#" + path[month+1:]
	}
	return path
}

func indexPerma(t *Tree, lang string) string {
	switch l := t.Level(); {
	case l == 0:
		return "/index"
	case l < 2:
		// prints /index/welt
		return t.Path(lang)
	}
	// prints /index/welt/wuestenleben-36c35dcb
	return fmt.Sprintf("%v-%v", t.Path(lang), t.Hash())
}

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
