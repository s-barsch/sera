package tree

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/read"
	"sacer/go/entry/tools/sort"
	"sacer/go/entry/types"
	"sacer/go/entry/types/set"
)

func readEntries(path string, parent *Tree) (entry.Entries, error) {
	fnErr := &tools.Err{
		Path: path,
		Func: "readEntries",
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

func readEntryFiles(files []*read.FileInfo, parent *Tree) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, fi := range files {
		if skipEntry(fi, parent) {
			continue
		}
		e, err := newEntry(fi.Path, parent)
		if err != nil {
			println(parent.Level())
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
		case "graph", "kine":
			if isDateTree(fi.Path, parent) {
				return true
			}
		case "index", "about", "extra":
			return true
		}
	}
	return false
}
