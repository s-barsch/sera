package tree

import (
	"fmt"
	"os"
	p "path/filepath"

	"g.rg-s.com/sacer/go/entry"
	he "g.rg-s.com/sacer/go/entry/tools"
	"g.rg-s.com/sacer/go/entry/tools/sort"
)

func readTrees(path string, parent *Tree) (Trees, error) {
	dirs, err := getTreeDirs(path, parent)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "readTrees",
			Err:  err,
		}
	}

	trees, err := readTreeDirs(dirs, parent)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "readTrees",
			Err:  err,
		}
	}

	entries, err := sort.SortEntries(path, toEntries(trees))
	if err != nil {
		return nil, err
	}

	return toTrees(entries)
}

// TODO: I did that – apparently – to use the same sort function I use for entries.
func toEntries(trees Trees) entry.Entries {
	es := entry.Entries{}
	for _, t := range trees {
		es = append(es, t)
	}
	return es
}

func toTrees(es entry.Entries) (Trees, error) {
	trees := Trees{}
	for _, e := range es {
		t, ok := e.(*Tree)
		if !ok {
			return nil, fmt.Errorf("toTrees: could not convert to *Tree")
		}
		trees = append(trees, t)
	}
	return trees, nil
}

func readTreeDirs(dirs []string, parent *Tree) (Trees, error) {
	trees := []*Tree{}
	for _, dirpath := range dirs {
		t, err := ReadTree(dirpath, parent)
		if err != nil {
			return nil, &he.Err{
				Path: dirpath,
				Func: "readTreeDirs",
				Err:  err,
			}
		}
		trees = append(trees, t)
	}
	return trees, nil
}

func getTreeDirs(path string, parent *Tree) ([]string, error) {
	l, err := os.ReadDir(path)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "getTreeDirs",
			Err:  err,
		}
	}

	dirs := []string{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())

		if he.IsDontIndex(fi.Name()) {
			continue
		}

		if !fi.IsDir() {
			continue
		}

		switch parent.Section() {
		case "graph", "cache":
			if !isDateTree(filepath, parent) {
				continue
			}
		}

		dirs = append(dirs, filepath)
	}
	return dirs, nil
}

func isDateTree(path string, parent *Tree) bool {
	return parent.Level() < 2 && he.FileType(path) == "dir"
}
