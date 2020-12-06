package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	p "path/filepath"
	"sacer/go/entry"
	he "sacer/go/entry/tools"
	"sacer/go/entry/tools/sort"
	//"sacer/go/entry/info"
	//"strings"
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

// I did that – apparently – to use the same sort function I use for entries.
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

func isSymLink(fi os.FileInfo) bool {
	return !(fi.Mode()&os.ModeSymlink != 0)
}

func getTreeDirs(path string, parent *Tree) ([]string, error) {
	l, err := ioutil.ReadDir(path)
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
		case "graph", "kine":
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

/*
	if !isSymLink(fi) && !fi.IsDir() {
		if !fi.IsDir() {
			continue
		}
	}
*/

/*
	// Holds that are completely empty are ommited.
	if len(h.Holds) == 0 && len(h.Entries) == 0 {
		continue
	}
*/

/*
func IsHold(path string) bool {
	info, err := info.ReadInfo(path)
	if err != nil {
		if strings.Contains(path, "/graph") {
			return true
		}
		return false
	}
	if info["inline"] == "true" {
		return false
	}
	return true
}
*/
