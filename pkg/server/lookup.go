package server

import (
	//"fmt"
	"stferal/pkg/entry"
)

func findElement(els entry.Els, acronym string) (interface{}, error) {
	id, err := entry.DecodeAcronym(acronym)
	if err != nil {
		return nil, err
	}
	return els.Lookup(id)
}

func findSet(sets entry.Sets, acronym string) (*entry.Set, error) {
	id, err := entry.DecodeAcronym(acronym)
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
	tree := &entry.Hold{}
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
