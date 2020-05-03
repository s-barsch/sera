package main

import (
	"io/ioutil"
	"sort"
)

func readTypes() ([]string, error) {
	types, err := readTypesDir(typeDir)
	if err != nil {
		return nil, err
	}

	media, err := readTypesDir(typeDir + "/media")
	if err != nil {
		return nil, err
	}

	types = append(media, types...)

	return types, nil
}

func readTypesDir(path string) ([]string, error) {
	types := []string{}
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fi := range l {
		if !fi.IsDir() || omitName(fi.Name()) {
			continue
		}
		types = append(types, fi.Name())
	}

	sort.Strings(types)

	return types, nil
}

func omitName(name string) bool {
	switch name {
	case ".wait", "media":
		return true
	}
	return false
}
