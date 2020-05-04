package tree

import (
	"fmt"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
	p "path/filepath"
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

func isGraph(path string, parent *Tree) bool {
	if parent == nil {
		return p.Base(path) == "graph"
	}
	return parent.Section() == "graph"
}

func readTreeInfo(path string, parent *Tree) (info.Info, error) {
	if !isGraph(path, parent) {
		return info.ReadDirInfo(path)
	}
	return readGraphInfo(path, parent)
}

func readGraphInfo(path string, parent *Tree) (info.Info, error) {
	date, err := parseGraphDate(path, parent)
	if err != nil {
		return nil, err
	}

	// if not present, we use the empty object.
	i, _ := info.ReadDirInfo(path)

	i["date"] = date.Format(helper.Timestamp)

	if parent == nil {
		return i, nil
	}

	switch parent.Level() {
	case 0:
		setBothLang(i, "title", date.Format("2006"))
		setBothLang(i, "label", date.Format("06"))
	case 1:
		monthDe := helper.GermanMonths[date.Month()] // Januar
		monthEn := date.Format("January")
		i["title"] = monthDe
		i["title-en"] = monthEn
		i["label"] = helper.Abbr(monthDe)
		i["label-en"] = helper.Abbr(monthEn)
		setBothLang(i, "slug", date.Format("01"))
	}
	return i, nil
}

func setBothLang(i info.Info, key, value string) {
	i[key] = value
	i[key + "-en"] = value
}

func parseGraphDate(path string, parent *Tree) (time.Time, error) {
	if parent == nil {
		return time.Parse("2006_01_02", "1991_01_02")
	}
	dirName := p.Base(path)
	switch parent.Level() {
	case 0:
		return time.Parse("06", dirName)
	case 1:
		t, err := time.Parse("06-01", dirName)
		if err != nil {
			return t, err
		}
		// Workaround so 2005 and 2005-01 won’t collide.
		if t.Month() == 1 {
			t = t.Add(time.Second)
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("Could not determine graph tree date. %v", path)
}
