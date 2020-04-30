package entry

import (
	"stferal/go/entry/file"
	"stferal/go/entry/helper"
	"stferal/go/entry/info"
	"time"
)

type Struct struct {
	Parent *Struct
	File   *file.File

	Date time.Time
	Info info.Info

	Entries []*Entry
	Structs Structs
}

type Structs []*Struct

func ReadStructure(path string, parent *Struct) (*Struct, error) {
	file, err := file.NewFile(path)
	if err != nil {
		return nil, err
	}

	inf, err := info.ReadInfo(path)
	if err != nil {
		return nil, err
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		return nil, helper.DateErr(path, err)
	}

	stru := &Struct{
		Parent: parent,
		File:   file,

		Date: date,
		Info: inf,
	}

	return stru, nil
}
