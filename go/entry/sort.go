package entry

import (
	"sort"
)

func (es Entries) SortAsc() {
	sort.Sort(Asc(es))
}

func (es Entries) SortDesc() {
	sort.Sort(Desc(es))
}

type Asc Entries

func (a Asc) Len() int { return len(a) }
func (a Asc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Asc) Less(i, j int) bool {
	return a[i].Id() < a[j].Id()
}

type Desc Entries

func (a Desc) Len() int { return len(a) }
func (a Desc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Desc) Less(i, j int) bool {
	return a[i].Id() > a[j].Id()
}


