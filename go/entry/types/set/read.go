package set

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media"
)

func readEntries(path string, parent entry.Entry) (entry.Entries, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "readEntries",
	}

	files, err := helper.GetFiles(path, false)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	reducedFiles := []string{}
	for _, f := range files {
		switch helper.FileType(f) {
		case "audio", "video", "html":
			continue
		}
		reducedFiles = append(reducedFiles, f)
	}

	entries, err := helper.ReadEntries(reducedFiles, parent, media.NewMediaEntry)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	// TODO: sorting

	return entries, nil
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
