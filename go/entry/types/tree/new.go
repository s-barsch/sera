package tree

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

type Tree struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	Entries entry.Entries
	Trees Trees
}

type Trees []*Tree

func ReadTree(path string, parent *Tree) (*Tree, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "ReadTree",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	inf, err := readTreeInfo(path, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseDate(inf["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Tree{
		parent: parent,
		file:   file,

		date: date,
		info: inf,
	}

	trees, err := readTrees(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.Entries = entries
	s.Trees = trees

	return s, nil
}

func readTreeInfo(path string, parent *Tree) (info.Info, error) {
	if section(path, parent) != "graph" {
		return info.ReadDirInfo(path)
	}
	return readGraphInfo(path, parent)
}


