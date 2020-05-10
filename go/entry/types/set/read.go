package set

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/helper/read"
	"stferal/go/entry/helper/sort"
	"stferal/go/entry/types/media"
)

func readEntries(path string, parent entry.Entry) (entry.Entries, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := read.GetFiles(path, false)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	entries, err := readEntryFiles(files, parent)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	return sort.SortEntries(path, entries)
}

func readEntryFiles(files []*read.FileInfo, parent entry.Entry) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, fi := range files {
		switch helper.FileType(fi.Path) {
		case "audio":
			continue
		}
		entry, err := media.NewMediaEntry(fi.Path, parent)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

/*
func readEntries(path string, parent *Tree) ([]*Entry, error) {

	els := []interface{}{}

	for _, fi := range l {

		fp := p.Join(path, fi.Name())


		if fi.Name() == "cache" {
			subels, err := readEls(p.Join(fp, "1600"), hold)
			if err != nil {
				return nil, err
			}
			els = append(els, subels...)
			continue
		}


		if fi.Name() == "cover.jpg" {
			// TODO
			continue
		}


		if helper.IsDontIndex(fp) {
			continue
		}


		e, err := NewEntry(fp, hold)
		if err != nil {
			return nil, err
		}


		els = append(els, e)
	}

	if hold.Info["private"] == "true" {
		l, err := makeElsPrivate(els)
		if err != nil {
			return nil, err
		}
		els = l
	}

	if exists(sortPath(path)) {
		sorted, err := SortEls(path, els)
		if err != nil {
			return nil, err
		}
		return sorted, nil
	}

		//if strings.Contains(path, "/graph/") {
	//		sort.Sort(Desc(els))
	//	} else {
	//	}

	sort.Sort(Asc(els))

	return els, nil
}

*/
