package server

import (
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
)

type SectionTree struct {
	Public, Private map[string]*tree.Tree
}

type SectionEntries struct {
	Public, Private map[string]entry.Entries
}

func (s *SectionTree) Local(local bool) map[string]*tree.Tree {
	if local {
		return s.Private
	}
	return s.Public
}

func (s *SectionEntries) Local(local bool) map[string]entry.Entries {
	if local {
		return s.Private
	}
	return s.Public
}

// trees

func (s *Server) LoadTrees() error {
	trees := map[string]*SectionTree{}
	recents := map[string]*SectionEntries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}

		trees[section] = &SectionTree{
			Private: makeLangs(t),
			Public:  makeLangs(t.Public()),
		}

		recents[section] = &SectionEntries{
			Private: serializeLangs(trees[section].Private),
			Public:  serializeLangs(trees[section].Public),
		}
	}

	s.Trees = trees
	s.Recents = recents

	return nil
}

func makeLangs(t *tree.Tree) map[string]*tree.Tree {
	return map[string]*tree.Tree{
		"de": t,
		"en": t.Lang("en"),
	}
}

func serializeLangs(langMap map[string]*tree.Tree) map[string]entry.Entries {
	return map[string]entry.Entries{
		"de": serialize(langMap["de"]),
		"en": serialize(langMap["en"]),
	}
}

func serialize(t *tree.Tree) entry.Entries {
	if t.Section() == "graph" {
		return t.TraverseEntriesReverse()
	}
	if t.Section() == "index" {
		ts := t.TraverseTrees()
		es := entry.Entries{}
		for _, e := range ts {
			es = append(es, e)
		}
		return es.Exclude().Desc()
	}
	return t.TraverseEntries().Exclude().Desc()
}
