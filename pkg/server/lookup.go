package server

import (
	//"fmt"
	"stferal/pkg/el"
)

func findElement(els el.Els, acronym string) (interface{}, error) {
	id, err := el.DecodeAcronym(acronym)
	if err != nil {
		return nil, err
	}
	return els.Lookup(id)
}

func findSet(sets el.Sets, acronym string) (*el.Set, error) {
	id, err := el.DecodeAcronym(acronym)
	if err != nil {
		return nil, err
	}
	return sets.Lookup(id)
}

/*
func (s *Server) lookupAcronymMulti(page, acronym string) (interface{}, error) {
	order := []string{}
	switch page {
	case "index":
		order = []string{"index", "graph"}
	case "graph":
		order = []string{"graph", "index"}
	case "about":
		order = []string{"about"}
	}
	for _, page := range order {
		e, err := s.LookupAcronym(page, acronym)
		if err != nil {
			//return nil, err
			continue
		}
		return e, nil
	}
	return nil, fmt.Errorf("lookupAcronymDual: Nothing found. %v", acronym)
}
*/

/*
func (s *Server) LookupAcronym(section, acronym string) (interface{}, error) {
	tree := &el.Hold{}
	switch page {
	case "index":
		tree = s.Trees["index"]
	case "graph":
		tree = s.Trees["graph"]
	case "about", "ueber":
		tree = s.Trees["about"]
	default:
		return nil, fmt.Errorf("lookupAcronym: unknown page. %v", page)
	}

	return tree.LookupAcronym(acronym)
}
*/
