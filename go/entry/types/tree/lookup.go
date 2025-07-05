package tree

import (
	"fmt"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
)

var ErrNilTree = fmt.Errorf("tree is nil")

func (t *Tree) LookupTreeHash(hash string) (*Tree, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	id, err := tools.ParseHash(hash)
	if err != nil {
		return nil, fmt.Errorf("fn LookupTreeHash: Couldn’t parse hash %v", err)
	}
	return t.LookupTree(id)
}

func (t *Tree) LookupTree(id int64) (*Tree, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	e, err := t.LookupEntry(id)
	if err != nil {
		return nil, err
	}
	tree, ok := e.(*Tree)
	if !ok {
		return nil, fmt.Errorf("entry with id %v (%v) found, but isn’t a tree", id, tools.ToTimestamp(id))
	}
	return tree, nil
}

func (t *Tree) LookupEntryHash(hash string) (entry.Entry, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	id, err := tools.ParseHash(hash)
	if err != nil {
		return nil, fmt.Errorf("fn LookupEntryHash: could not parse hash %v", err)
	}
	return t.LookupEntry(id)
}

// Starting recursive function
func (t *Tree) LookupEntry(id int64) (entry.Entry, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	return t.lookup([]*Tree{}, id)
}

// Recursive function
func (t *Tree) lookup(stack []*Tree, id int64) (entry.Entry, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	if t.Id() == id {
		return t, nil
	}
	for _, e := range t.Entries() {
		if e.Id() == id {
			return e, nil
		}
	}
	for i, h := range t.Trees {
		if i == 0 {
			return h.lookup(append(stack, t.Trees[1:]...), id)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.lookup(stack[1:], id)
		}
	}
	return nil, fmt.Errorf("fn lookupEntry: Id %v (%v) not found", id, tools.ToTimestamp(id))
}

// search

func (t *Tree) SearchTree(slug, lang string) (*Tree, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	return t.search([]*Tree{}, slug, lang)
}

func (t *Tree) search(stack []*Tree, slug, lang string) (*Tree, error) {
	if t == nil {
		return nil, ErrNilTree
	}

	if t.Slug(lang) == slug {
		return t, nil
	}
	//	for _, e := range t.Els {
	//		if e.Id() == id {
	//			return e, nil
	//		}
	//	}
	for i, h := range t.Trees {
		if i == 0 {
			return h.search(append(stack, t.Trees[1:]...), slug, lang)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.search(stack[1:], slug, lang)
		}
	}

	return nil, fmt.Errorf("could not find slug %v in Tree", slug)
}
