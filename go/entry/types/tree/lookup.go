package tree

import (
	"fmt"
	"stferal/go/entry/helper"
)

func (t *Tree) LookupHash(hash string) (interface{}, error) {
	id, err := helper.FromB16(hash)
	if err != nil {
		return nil, fmt.Errorf("LookupHash: %v", err)
	}
	return t.Lookup(id)
}

func (t *Tree) Lookup(id int64) (interface{}, error) {
	return t.lookup([]*Tree{}, id)
}

func (t *Tree) lookup(stack []*Tree, id int64) (interface{}, error) {
	if t.Id() == id {
		return t, nil
	}
	for _, e := range t.Entries {
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
	return nil, fmt.Errorf("Couldn’t find id %v in Tree.", id)
}

/*
func (t *Tree) Search(name, lang string) (*Tree, error) {
	return t.search([]*Tree{}, name, lang)
}

func (t *Tree) search(stack []*Tree, name, lang string) (*Tree, error) {
	if t.Name(lang) == name {
		return hold, nil
	}
	//	for _, e := range t.Els {
	//		if Id(e) == id {
	//			return e, nil
	//		}
	//	}
	for i, h := range t.Trees {
		if i == 0 {
			return h.search(append(stack, t.Trees[1:]...), name, lang)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.search(stack[1:], name, lang)
		}
	}
	return nil, fmt.Errorf("Couldn’t find name %v in Tree.", name)
}

*/
