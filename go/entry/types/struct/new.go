package stru

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

type Struct struct {
	Parent entry.Entry
	File   *file.File

	date time.Time
	info info.Info

	Entries entry.Entries
	Structs Structs
}

type Structs []*Struct

func ReadStruct(path string, parent entry.Entry) (*Struct, error) {
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

	s := &Struct{
		Parent: parent,
		File:   file,

		date: date,
		info: inf,
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	structs, err := readStructs(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.Entries = entries
	s.Structs = structs

	return s, nil
}
