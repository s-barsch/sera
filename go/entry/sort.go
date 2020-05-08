package entry

import (
	"sort"
)

func (es Entries) Asc() Entries {
	es.SortAsc()
	return es
}

func (es Entries) Desc() Entries {
	es.SortDesc()
	return es
}

func (es Entries) SortAsc() {
	sort.Sort(asc(es))
}

func (es Entries) SortDesc() {
	sort.Sort(desc(es))
}

type asc Entries

func (a asc) Len() int      { return len(a) }
func (a asc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a asc) Less(i, j int) bool {
	return a[i].Id() < a[j].Id()
}

type desc Entries

func (a desc) Len() int      { return len(a) }
func (a desc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a desc) Less(i, j int) bool {
	return a[i].Id() > a[j].Id()
}
