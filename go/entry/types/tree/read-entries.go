package tree

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/helper/read"
	"stferal/go/entry/types/media"
	"stferal/go/entry/types/set"
)

func readEntries(path string, parent *Tree) (entry.Entries, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := read.GetFiles(path, true)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	reducedFiles := []string{}
	for _, f := range files {
		switch helper.FileType(f) {
		case "audio", "video", "html":
			continue
		}
		reducedFiles = append(reducedFiles, f)
	}

	entries, err := readEntriesLoop(reducedFiles, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	// TODO: sorting

	return entries, err
}

func readEntriesLoop(files []string, parent *Tree) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, path := range files {
		if graphTree(path, parent) {
			continue
		}
		e, err := newEntry(path, parent)
		if err != nil {
			println(parent.Level())
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func graphTree(path string, parent *Tree) bool {
	return parent.Level() < 3 && isGraph(path, parent) && helper.FileType(path) == "dir"
}

func newEntry(path string, parent *Tree) (entry.Entry, error) {
	switch helper.FileType(path) {
	case "file":
		break
	case "dir":
		return set.NewSet(path, parent)
	default:
		return media.NewMediaEntry(path, parent)
	}
	return nil, &helper.Err{
		Path: path,
		Func: "newObjFunc",
		Err:  fmt.Errorf("invalid entry type: %v", helper.FileType(path)),
	}
}
