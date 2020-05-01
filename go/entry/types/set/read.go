package entry

import (
	"io/ioutil"
	p "path/filepath"
	"stferal/go/entry/helper"
)

func ReadEntries(path string, parent *Struct) ([]*Entry, error) {
	files, err := getEntryFiles(path)
	if err != nil {
		return nil, err
	}

	entries, err := readEntryFiles(files, parent)
	if err != nil {
		return nil, err
	}

	// TODO: sorting

	return entries, err
}

func readEntryFiles(files []string, parent *Struct) ([]*Entry, error) {
	entries := Entries{}
	for _, filepath := range files {
		entry, err := NewEntry(filepath, parent)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func getEntryFiles(path string) ([]string, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	list := []string{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())

		if fi.Name() == "cache" {
			imageFolder := p.Join(filepath, "1600")
			images, err := getEntryFiles(imageFolder)
			if err != nil {
				return nil, err
			}
			list = append(list, images...)
			continue
		}

		if helper.IsDontIndex(filepath) {
			continue
		}

		list = append(list, filepath)
	}

	return list, nil
}

/*
func readEntries(path string, parent *Struct) ([]*Entry, error) {

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
