package tree

import (
	p "path/filepath"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"sacer/go/entry/types/text"
	"sacer/go/entry/types/set"
	"time"
)

type Tree struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	Trees   Trees

	Footnotes text.Footnotes
}

func (t *Tree) Copy() *Tree {
	return &Tree{
		parent: t.parent,
		file:   t.file.Copy(),

		date: t.date,
		info: t.info.Copy(),

		entries: t.entries,
		Trees:   t.Trees,
	}
}

type Trees []*Tree

func ReadTree(path string, parent *Tree) (*Tree, error) {
	fnErr := &tools.Err{
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

	date, err := tools.ParseTimestamp(inf["date"])
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

	s.entries = entries
	s.Trees = trees

	s.Footnotes = set.NumberFootnotes(s.Entries())

	return s, nil
}

func readTreeInfo(path string, parent *Tree) (info.Info, error) {
	if !isGraph(path, parent) {
		return info.ReadDirInfo(path)
	}
	return readGraphInfo(path, parent)
}

// Function only needed here, not in readTrees or readEntries.
// Because in these, #parent# will always be defined.
func isGraph(path string, parent *Tree) bool {
	if parent == nil {
		return isGraphSection(p.Base(path))
	}
	return isGraphSection(parent.Section())
}

func isGraphSection(section string) bool {
	switch section {
	case "graph", "kine":
		return true
	}
	return false
}
