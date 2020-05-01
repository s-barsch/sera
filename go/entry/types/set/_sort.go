package set

import (
	"sort"
	"path/filepath"
	"io/ioutil"
)

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
