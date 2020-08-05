package set

import (
	// "log"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"stferal/go/entry/types/media/image"
	"time"
)

type Set struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	Cover   *image.Image
}

func (s *Set) Copy() *Set {
	return &Set{
		parent: s.parent,
		file:   s.file,

		date: s.date,
		info: s.info,

		entries: s.entries,
		Cover: s.Cover,
	}
}

type Sets []*Set

func NewSet(path string, parent entry.Entry) (*Set, error) {
	fnErr := &helper.Err{
		Path: path,
		Func: "NewSet",
	}

	file, err := file.NewFile(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	info, err := info.ReadDirInfo(path)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	date, err := helper.ParseTimestamp(info["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Set{
		parent: parent,
		file:   file,

		date: date,
		info: info,
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.Cover, s.entries = extractCover(entries)
	// s.entries = entries

	return s, nil
}

func extractCover(es entry.Entries) (*image.Image, entry.Entries) {
	for i, e := range es {
		if e.File().Name() == "cover.jpg" {
			img, ok := e.(*image.Image)
			if ok {
				return img, append(es[:i], es[i+1:]...)
			}
		}
	}
	return nil, es
}

/*
	cover, err := ReadCover(path, h)
	if err != nil {
		// log.Println(err)
	}
*/
