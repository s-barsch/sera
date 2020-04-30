package main

import (
	"io/ioutil"
	"sort"
)

type Types struct {
	media []string
	dirs  []string
	file  string
}

func (t *Types) All() []string {
	return append(t.NoFile(), t.file)
}

func (t *Types) Media() []string {
	return t.media
}

func (t *Types) NoFile() []string {
	return append(t.media, t.dirs...)
}

func readTypes() (*Types, error) {
	types, err := readTypesDir()
	if err != nil {
		return nil, err
	}

	return split(types), nil
}

func readTypesDir() ([]string, error) {
	types := []string{}
	l, err := ioutil.ReadDir("./types")
	if err != nil {
		return nil, err
	}

	for _, fi := range l {
		types = append(types, fi.Name())
	}

	return types, nil
}

func split(types []string) *Types {
	dirs, media := []string{}, []string{}

	for _, t := range types {
		switch t {
		case "hold", "set":
			dirs = append(dirs, t)
		case "file":
			continue
		default:
			media = append(media, t)
		}
	}

	// asc media
	sort.Strings(media)

	// desc dirs
	sort.Sort(sort.Reverse(sort.StringSlice(dirs)))

	return &Types{
		media: media,
		dirs:  dirs,
		file:  "file",
	}
}
