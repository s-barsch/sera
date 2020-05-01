package stru

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media"
	"stferal/go/entry/types/set"
)

func readEntries(path string, parent interface{}) ([]*entry.Entry, error) {
	files, err := helper.GetFiles(path, true)
	if err != nil {
		return nil, err
	}

	entries, err := helper.ReadEntries(files, parent, newObjFunc)
	if err != nil {
		return nil, err
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
		return media.NewMedia(path)
	}
	return nil, helper.TypeErr(path)
}
