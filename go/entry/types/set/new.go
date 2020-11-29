package set

import (
	// "log"
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/file"
	"sacer/go/entry/info"
	"sacer/go/entry/types/image"
	"sacer/go/entry/types/text"
	"time"
)

type Set struct {
	parent entry.Entry
	file   *file.File

	date time.Time
	info info.Info

	entries entry.Entries
	Cover   *image.Image

	Kine      entry.Entries
	Notes     entry.Entries

	Footnotes text.Footnotes
}

func (s *Set) Copy() *Set {
	return &Set{
		parent: s.parent,
		file:   s.file,

		date: s.date,
		info: s.info.Copy(),

		entries: s.entries,
		Cover: s.Cover,

		Notes: s.Notes,
		Kine:  s.Kine,
		
		Footnotes: s.Footnotes,
	}
}

type Sets []*Set

func NewSet(path string, parent entry.Entry) (*Set, error) {
	fnErr := &tools.Err{
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

	date, err := tools.ParseTimestamp(info["date"])
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

	s.Notes, s.entries = extractNotes(s.entries)

	s.Footnotes = NumberFootnotes(s.Entries())

	return s, nil
}

func extractNotes(es entry.Entries) (entry.Entries, entry.Entries) {
	notes := entry.Entries{}
	entries := entry.Entries{}
	for _, e := range es {
		if e.Info().Note() == true {
			notes = append(notes, e)
			continue
		}
		entries = append(entries, e)
	}
	return notes, entries
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
