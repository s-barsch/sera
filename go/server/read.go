package server

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/tree"
	"sort"
)

var sections = []string{
	"kine",
	"indecs",
	"graph",
	"about",
	"extra",
	"log",
}

type DoubleTree struct {
	All, Blur map[string]*tree.Tree
}

type DoubleEntries struct {
	All, Blur map[string]entry.Entries
}

func (d *DoubleTree) Access(subscriber bool) map[string]*tree.Tree {
	if subscriber {
		return d.All
	}
	return d.Blur
}

func (d *DoubleEntries) Access(subscriber bool) map[string]entry.Entries {
	if subscriber {
		return d.All
	}
	return d.Blur
}

func (s *Server) ReadTrees() error {
	trees := map[string]*DoubleTree{}
	recents := map[string]*DoubleEntries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}

		if !s.Flags.Local {
			t = t.Public()
		}

		trees[section] = &DoubleTree{
			All:  makeLangs(t),
			Blur: makeLangs(t.Blur()),
		}

		recents[section] = &DoubleEntries{
			All:  serializeLangs(trees[section].All),
			Blur: serializeLangs(trees[section].Blur),
		}

	}

	s.Trees = trees
	s.Recents = recents

	return nil
}

func makeLangs(t *tree.Tree) map[string]*tree.Tree {
	return map[string]*tree.Tree{
		"de": t,
		"en": t.Translated("en"),
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
	case "indecs":
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
