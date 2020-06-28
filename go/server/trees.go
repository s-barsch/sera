package server

import (
	"fmt"
	"sort"
	"stferal/go/entry"
	"stferal/go/entry/helper"
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
		for _, te := range ts {
			if len(te.Entries()) == 0 {
				continue
			}
			es = append(es, te)
		}
		es = es.Exclude()
		sort.Sort(byRevision(es))
		return es
	}
	return t.TraverseEntries().Exclude().Desc()
}

type byRevision entry.Entries

func (a byRevision) Len() int      { return len(a) }
func (a byRevision) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a byRevision) Less(i, j int) bool {
	return newestDate(a[i]) > newestDate(a[j])
}

func newestDate(e entry.Entry) int64 {
	if rev := e.Info()["revision"]; rev != "" {
		t, err := helper.ParseTimestamp(rev)
		if err != nil {
			fmt.Println(err)
			return e.Id()
		}
		return t.Unix()
	}
	return e.Id()
}
