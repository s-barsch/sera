package set

import (
	// "log"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"stferal/go/entry/parts/file"
	"time"
)

type Set struct {
	File  *file.File

	Date  time.Time
	Info  info.Info

	//Cover *Image
	//Els Els
}

type Sets []*Set

func NewSet(path string, parent interface{}) (*Set, error) {
	file, err := file.New(path)
	if err != nil {
		return nil, err
	}

	info, err := info.ReadInfo(path)
	if err != nil {
		return nil, err
	}

	date, err := helper.ParseDate(info["date"])
	if err != nil {
		return nil, helper.DateErr(path, err)
	}

	/*
	// sketchy
	h := &Hold{
		Parent: parent,
		File:   file,

		Date: date,
		Info: info,
	}

	els, err := readEls(path, h)
	if err != nil {
		return nil, err
	}

	cover, err := ReadCover(path, h)
	if err != nil {
		// log.Println(err)
	}
	*/

	s := &Set{
		File:   file,

		Date:  date,
		Info:  info,

		//Cover: cover,
		//Els: els,
	}

	return s, nil
}
