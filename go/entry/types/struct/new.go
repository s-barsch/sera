package stru

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

type Struct struct {
	Parent *Struct
	File   *file.File

	Date time.Time
	Info info.Info

	Entries entry.Entries
	Structs Structs
}

type Structs []*Struct

func ReadStruct(path string, parent *Struct) (*Struct, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "ReadStruct",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	// TODO: Graph needs a specific way
	inf, err := info.ReadDirInfo(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntries(path, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	structs, err := readStructs(path, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	stru := &Struct{
		Parent: parent,
		File:   file,

		Date: date,
		Info: inf,

		Entries: entries,
		Structs: structs,
	}

	return stru, nil
}
