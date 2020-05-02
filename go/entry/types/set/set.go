package set

import (
	// "log"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

type Set struct {
	File *file.File

	Date time.Time
	Info info.Info

	Entries entry.Entries
	//Cover *Image
}

type Sets []*Set

func NewSet(path string) (*Set, error) {
	file, err := file.New(path)
	if err != nil {
		return nil, err
	}

	info, err := info.ReadInfoDir(path)
	if err != nil {
		return nil, err
	}

	date, err := helper.ParseDate(info["date"])
	if err != nil {
		return nil, helper.DateErr(path, err)
	}

	s := &Set{
		File: file,

		Date: date,
		Info: info,
	}

	entries, err := readEntries(path, s)
	if err != nil {
		return nil, err
	}

	s.Entries = entries

	/*
		cover, err := ReadCover(path, h)
		if err != nil {
			// log.Println(err)
		}
	*/

	return s, nil
}
