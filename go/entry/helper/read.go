package helper

import (
	"stferal/go/entry"
	"io/ioutil"
	p "path/filepath"
)

func ReadEntries(paths []string, parent interface{}, newObjFunc entry.NewObjFunc) (entry.Entries, error) {
	entries := entry.Entries{}
	for _, path := range paths {
		entry, err := entry.NewEntry(path, parent, newObjFunc)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func GetFiles(path string, withDirs bool) ([]string, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	list := []string{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())

		if fi.Name() == "cache" {
			imageFolder := p.Join(filepath, "1600")
			images, err := GetFiles(imageFolder, false)
			if err != nil {
				return nil, err
			}
			list = append(list, images...)
			continue
		}

		if !withDirs && fi.IsDir() {
			continue
		}

		if IsDontIndex(filepath) {
			continue
		}

		list = append(list, filepath)
	}

	return list, nil
}

