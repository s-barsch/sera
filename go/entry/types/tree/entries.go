package tree

import (
	"fmt"
	"os"
	p "path/filepath"
	"regexp"
	gosort "sort"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/tools/read"
	"g.rg-s.com/sera/go/entry/tools/sort"
	media "g.rg-s.com/sera/go/entry/types"
	"g.rg-s.com/sera/go/entry/types/set"
)

// isMergeTree checks if folder name "06-0102" is present
func isMergeTree(path string) bool {
	return regexp.MustCompile(`\d{2}-\d{4}`).MatchString(p.Base(path))
}

func readEntries(path string, parent *Tree) (entry.Entries, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "readEntries",
	}

	if isMergeTree(path) {
		return readMergeTreeEntries(path, parent)
	}

	files, err := read.GetFiles(path, true)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntryFiles(files, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return sort.SortEntries(path, entries)
}

func readMergeTreeEntries(path string, parent *Tree) (entry.Entries, error) {
	l, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	dirsAsc(l)
	entries := entry.Entries{}
	for _, dir := range l {
		if !dir.IsDir() {
			continue
		}
		es, err := readEntries(p.Join(path, dir.Name()), parent)
		if err != nil {
			return entries, err
		}
		entries = append(entries, es...)
	}
	return sort.SortEntries(path, entries)
}

func dirsAsc(files []os.DirEntry) {
	gosort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
}

func readEntryFiles(files []*read.FileInfo, parent *Tree) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, fi := range files {
		if skipEntry(fi, parent) {
			continue
		}
		e, err := newEntry(fi.Path, parent)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func newEntry(path string, parent *Tree) (entry.Entry, error) {
	switch tools.FileType(path) {
	case "file":
		break
	case "dir":
		return set.NewSet(path, parent)
	default:
		return media.NewMediaEntry(path, parent)
	}
	return nil, &tools.Err{
		Path: path,
		Func: "newObjFunc",
		Err:  fmt.Errorf("invalid entry type: %v", tools.FileType(path)),
	}
}

func skipEntry(fi *read.FileInfo, parent *Tree) bool {
	if fi.IsDir() {
		switch parent.Section() {
		case "graph", "cache":
			if isDateTree(fi.Path, parent) {
				return true
			}
		case "indecs", "about", "extra":
			return true
		}
	}
	return false
}
