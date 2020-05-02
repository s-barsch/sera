package stru

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media"
	"stferal/go/entry/types/set"
)

func readEntries(path string, parent interface{}) (entry.Entries, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := helper.GetFiles(path, true)
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

	entries, err := helper.ReadEntries(reducedFiles, parent, newObjFunc)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	// TODO: sorting

	return entries, err
}

func newObjFunc(path string, parent interface{}) (entry.Entry, error) {
	switch helper.FileType(path) {
	case "file":
		break
	case "dir":
		return set.NewSet(path)
	default:
		return media.NewMediaObj(path, parent)
	}
	return nil, &helper.Err{
		Path: path,
		Func: "newObjFunc",
		Err:  fmt.Errorf("invalid entry type: %v", helper.FileType(path)),
	}
}
