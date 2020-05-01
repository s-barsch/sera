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
	file, err := file.New(path)
	if err != nil {
		return nil, err
	}

	// TODO: Graph needs a specific way
	inf, err := info.ReadInfo(path)
	if err != nil {
		return nil, err
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		return nil, helper.DateErr(path, err)
	}

	entries, err := readEntries(path, parent)
	if err != nil {
		return nil, err
	}

	structs, err := readStructs(path, parent)
	if err != nil {
		return nil, err
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
