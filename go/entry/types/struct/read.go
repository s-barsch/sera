package stru

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media"
	"stferal/go/entry/types/set"
)

func readEntries(path string, parent interface{}) ([]*entry.Entry, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := helper.GetFiles(path, true)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := helper.ReadEntries(files, parent, newObjFunc)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	// TODO: sorting

	return entries, err
}

func newObjFunc(path string) (interface{}, error) {
	switch helper.FileType(path) {
	case "file":
		break
	case "dir":
		return set.NewSet(path)
	default:
		return media.NewMediaObj(path)
	}
	return nil, &helper.Err{
		Path: path,
		Func: "newObjFunc",
		Err:  fmt.Errorf("invalid entry type: %v", helper.FileType(path)),
	}
}
