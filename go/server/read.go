package server

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sort"
)

type langTrees   map[string]*tree.Tree
type langEntries map[string]entry.Entries

func (s *Server) ReadTrees() error {
	trees := map[string]langTrees{}
	recents := map[string]langEntries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}

		if !s.Flags.Local {
			t = t.Public()
		}

		trees[section] = map[string]*tree.Tree{
			"de": t,
			"en": t.Lang("en"),
		}

		recents[section] = map[string]entry.Entries{
			"de": serialize(trees[section]["de"]),
			"en": serialize(trees[section]["en"]),
		}
	}
		
	s.Trees = trees
	s.Recents = recents

	return nil
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
