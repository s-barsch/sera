package entry

import (
	"fmt"
	"stferal/go/entry/file"
	"stferal/go/entry/info"
	"time"
)

type Struct struct {
	Parent  *Struct
	File    *file.File

	Date    time.Time
	Info    info.Info

	Entries  []*Entry
	Structs  Structs
}

type Structs []*Struct

func ReadStructure(path string, parent *Struct) (*Struct, error) {
	file, err := file.NewFile(path)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("not implemented")
}
