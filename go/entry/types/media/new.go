package media

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media/text"
	"stferal/go/entry/types/media/image"
)

func NewMediaObj(path string, parent interface{}) (entry.Entry, error) {
	switch helper.FileType(path) {
	case "text":
		return text.NewText(path)
	case "image":
		return image.NewImage(path)
	}
	return nil, &helper.Err{
		Path: path,
		Func: "NewMediaObj",
		Err:  fmt.Errorf("invalid entry type: %v", helper.FileType(path)),
	}
}
