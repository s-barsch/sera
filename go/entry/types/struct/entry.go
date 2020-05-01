package stru

import (
	"fmt"
	"stferal/go/entry"
)

func newStructEntry(path string, parent interface{}) (*entry.Entry, error) {
	switch helper.EntryGroup(path) {
	case "media":
		return media.NewMedia(filepath, parent)
	case "structure":
		return set.NewSet(filepath, parent)
	}
	return nil, fmt.Errorf("invalid entry file")
}
