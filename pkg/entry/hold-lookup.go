package entry

import (
	"fmt"
)

type Holds []*Hold

func (holds Holds) Reverse() Holds {
	n := Holds{}
	for i := len(holds) - 1; i >= 0; i-- {
		n = append(n, holds[i])
	}
	return n
}

/*
func (hold *Hold) TraverseEls() Els {
	els := hold.traverseEls([]*Hold{})
	//sort.Sort(Desc(els))
	return els
}
*/

func (hold *Hold) TraverseElsReverse() Els {
	holds := hold.TraverseHolds()

	els := Els{}

	for _, h := range holds {
		els = append(els, h.Els.Reverse()...)
	}

	return els
}

func newEls(els Els) Els {
	nels := Els{}
	for _, e := range els {
		nels = append(nels, e)
	}
	return nels
}

func (hold *Hold) TraverseEls() Els {
	holds := hold.TraverseHolds()

	els := Els{}

	for _, h := range holds {
		els = append(els, h.Els...)
		//sort.Sort(Desc(h.Els))
	}

	return els
}

func (hold *Hold) TraverseHolds() Holds {
	holds := Holds{hold}
	for _, h := range hold.Holds.Reverse() {
		hs := h.TraverseHolds()
		holds = append(holds, hs...)
	}
	return holds
}

/*
func (hold *Hold) traverseEls(stack []*Hold) Els {
	els := Els{}
	for _, e := range hold.Els {
		els = append(els, e)
	}
	for i, h := range hold.Holds.Reverse() {
		if i == 0 {
			return append(els, h.traverseEls(append(stack, hold.Holds[1:]...))...)
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

func (hold *Hold) LookupAcronym(acronym string) (interface{}, error) {
	id, err := DecodeAcronym(acronym)
	if err != nil {
		return nil, fmt.Errorf("LookupAcronym: %v", err)
	}
	return hold.Lookup(id)
}

func (hold *Hold) Search(name, lang string) (*Hold, error) {
	return hold.search([]*Hold{}, name, lang)
}

func (hold *Hold) search(stack []*Hold, name, lang string) (*Hold, error) {
	if hold.Name(lang) == name {
		return hold, nil
	}
	/*
		for _, e := range hold.Els {
			if Id(e) == id {
				return e, nil
			}
		}
	*/
	for i, h := range hold.Holds {
		if i == 0 {
			return h.search(append(stack, hold.Holds[1:]...), name, lang)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.search(stack[1:], name, lang)
		}
	}
	return nil, fmt.Errorf("Couldn’t find name %v in Tree.", name)
}

func (hold *Hold) Lookup(id string) (interface{}, error) {
	return hold.lookup([]*Hold{}, id)
}

func (hold *Hold) lookup(stack []*Hold, id string) (interface{}, error) {
	if hold.Id() == id {
		return hold, nil
	}
	for _, e := range hold.Els {
		if Id(e) == id {
			return e, nil
		}
	}
	for i, h := range hold.Holds {
		if i == 0 {
			return h.lookup(append(stack, hold.Holds[1:]...), id)
		}
	}
	for i, h := range stack {
		if i == 0 {
			return h.lookup(stack[1:], id)
		}
	}
	return nil, fmt.Errorf("Couldn’t find id %v in Tree.", id)
}
