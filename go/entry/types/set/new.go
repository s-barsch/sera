package entry

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	// "log"
)

type Set struct {
	File *File

	Date  time.Time
	Info  Info
	Cover *Image

	Els Els
}

type Sets []*Set

/*
//func markParent(els Els, acronym string) {
func markParent(els Els, acronym string) {
	for _, e := range els {
		switch e.(type) {
		case *Image:
			e.(*Image).File.Parent = acronym
		case *File:
			e.(*File).Parent = acronym
		}
	}
	return
}
*/

func NewSet(path string, mother *Hold) (*Set, error) {
	file, err := NewFile(path, mother)
	if err != nil {
		return nil, err
	}

	info, err := ReadInfo(path)
	if err != nil {
		return nil, err
	}

	date, err := ParseDate(info["date"])
	if err != nil {
		return nil, invalidDate(path, err)
	}

	file.Id = info["date"]

	h := &Hold{
		Mother: mother,
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

	s := &Set{
		File: file,

		Date:  date,
		Info:  info,
		Cover: cover,

		Els: els,
	}

	return s, nil
}

type SetDesc []*Set

func (a SetDesc) Len() int           { return len(a) }
func (a SetDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SetDesc) Less(i, j int) bool { return a[i].File.Id > a[j].File.Id }

func sortPath(path string) string {
	return filepath.Join(path, "/.sort")
}

func SortEls(path string, els Els) (Els, error) {
	b, err := ioutil.ReadFile(sortPath(path))
	if err != nil {
		return nil, err
	}
	l := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, sortElement := range reverse(l) {
		//for _, sortElement := range l {
		for i, e := range els {
			f, err := ElFileSafe(e)
			if err != nil {
				return els, fmt.Errorf("SortEls: %v", err)
			}
			if filepath.Base(f.Path) == sortElement {
				cut := els[i]
				els = append(Els{cut}, append(els[:i], els[i+1:]...)...)
			}
		}
	}
	return els, nil
}

func reverse(ss []string) []string {
	ns := []string{}
	for i := len(ss) - 1; i >= 0; i-- {
		ns = append(ns, ss[i])
	}
	return ns
}

/*
func getState(state string) string {
	if state != "" {
		return strings.ToLower(state)
	}
	return "live"
}
*/
