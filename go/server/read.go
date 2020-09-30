package server

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sort"
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
			Private: filterLangs(t),
			Public:  filterLangs(t.Public()),
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

func filterLangs(t *tree.Tree) map[string]*tree.Tree {
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
	switch t.Section() {
	case "graph", "kine":
		return t.TraverseEntriesReverse()
	case "index":
		es := entry.Entries{}
		for _, tree := range t.TraverseTrees() {
			if len(tree.Entries()) == 0 {
				continue
			}
			es = append(es, tree)
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
		t, err := tools.ParseTimestamp(rev)
		if err != nil {
			fmt.Println(err)
			return e.Id()
		}
		return t.Unix()
	}
	return e.Id()
}