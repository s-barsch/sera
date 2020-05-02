package media

import (
	"fmt"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media/text"
)

func NewMediaObj(path string) (interface{}, error) {
	switch helper.FileType(path) {
	case "text":
		return text.NewText(path)
	}
	return nil, &helper.Err{
		Path: path,
		Func: "NewMediaObj",
		Err:  fmt.Errorf("invalid entry type: %v", helper.FileType(path)),
	}
}
