package set

import (
	// "log"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (s *Set) Id() string {
	return "sample"
}

type Set struct {
	File *file.File

	Date time.Time
	Info info.Info

	Entries entry.Entries
	//Cover *Image
}

type Sets []*Set

func NewSet(path string) (*Set, error) {
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

	date, err := helper.ParseDate(info["date"])
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s := &Set{
		File: file,

		Date: date,
		Info: info,
	}

	entries, err := readEntries(path, s)
	if err != nil {
		fnErr.Err = err
		return nil, fnErr
	}

	s.Entries = entries

	return s, nil
}

	/*
		cover, err := ReadCover(path, h)
		if err != nil {
			// log.Println(err)
		}
	*/

