package entry

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	// "log"
)

type Sets []*Set

/*
type Set struct {
	File *File

	Page string

	Date  time.Time
	Info  Info
	State string
	Cover *Image

	Els Els
}
*/

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

func ReadCover(path string, hold *Hold) (*Image, error) {
	return NewImage(filepath.Join(path, "cover.jpg"), hold)
}

func readEls(path string, hold *Hold) (Els, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	els := Els{}
	for _, fi := range l {
		fp := filepath.Join(path, fi.Name())
		if fi.Name() == "cache" {
			subels, err := readEls(filepath.Join(fp, "1600"), hold)
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
		if isDontIndex(fp) {
			continue
		}
		e, err := NewEl(fp, hold)
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

	/*
		if strings.Contains(path, "/graph/") {
			sort.Sort(Desc(els))
		} else {
		}
	*/
	sort.Sort(Asc(els))

	return els, nil
}

func makeElsPrivate(els Els) (Els, error) {
	l := Els{}
	for _, e := range els {
		i, err := EntryInfo(e)
		if err != nil {
			return els, err
		}
		i["private"] = "true"
		err = setInfo(e, i)
		if err != nil {
			return els, err
		}
		l = append(l, e)
	}
	return l, nil
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
