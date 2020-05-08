package entry

import (
	"fmt"
	"sort"
	"time"
)

type Els []interface{}

type Args []*Arg

type Arg struct {
	El   interface{}
	Lang string
}

func (args Args) Offset(start, end int) Args {
	l := len(args)
	if l < start {
		return Args{}
	}
	if end > l || end <= 0 {
		return args[start:]
	}
	return args[start:end]
}

func (els Els) Limit(n int) Els {
	if len(els) <= n {
		return els
	}
	return els[:n]
}

func (els Els) Args(lang string) Args {
	args := Args{}
	for _, e := range els {
		args = append(args, &Arg{
			El:   e,
			Lang: lang,
		})
	}
	return args
}

func (els Els) Year(year int) Els {
	nl := Els{}

	for _, e := range els {
		if Date(e).Year() == year {
			nl = append(nl, e)
			continue
		}
		break
	}
	return nl
}

/*
func (els Els) LazyLoad() Els {
	imgs := 0
	nl := Els{}
	for _, e := range els {
		i, ok := e.(*Image)
		if ok {
			imgs++
			if imgs >= 12 {
				ni := i
				ni.LazyLoad = true
				nl = append(nl, ni)
				continue
			}
		}
		nl = append(nl, e)
	}
	return nl
}
*/

/*
type Desc Els

func (a Desc) Len() int {
	return len(a)
}

func (a Desc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Desc) Less(i, j int) bool {
	return Id(a[i]) > Id(a[j])
}
*/

func (sets Sets) Lookup(id string) (*Set, error) {
	for _, s := range sets {
		if s.File.Id != id {
			continue
		}
		return s, nil
	}
	return nil, NotFoundErr(id)
}

func (els Els) LookupPosition(id string) (int, error) {
	for i, e := range els {
		if Id(e) != id {
			continue
		}
		return i, nil
	}
	return -1, NotFoundErr(id)
}

func (els Els) Lookup(id string) (interface{}, error) {
	for _, e := range els {
		if Id(e) != id {
			continue
		}
		return e, nil
	}
	return nil, NotFoundErr(id)
}

/*
func (els Els) FindImage(filename string) (*Image, error) {
	println(filename)
	for _, e := range els {
		switch e.(type) {
		case *Image:
			if e.(*Image).File.Base() == filename {
				return e.(*Image), nil
			}
		}
	}
	return nil, fmt.Errorf("Couldnt find image %v in given Els.", filename)
}
*/

func NotFoundErr(id string) error {
	return fmt.Errorf("Lookup: el not found: %v", id)
}

type Day struct {
	Date time.Time
	Els  Els
}

func (els Els) Asc() Els {
	sort.Sort(Asc(els))
	return els
}

func (els Els) Desc() Els {
	sort.Sort(Desc(els))
	return els
}

func (els Els) Days() []Els {
	days := []Els{}
	cd := Els{}
	for i, e := range els {
		if i == 0 || !isNewDay(e, els[i-1]) {
			cd = append(cd, e)
			continue
		}
		days = append(days, cd.Reverse())
		cd = Els{e}
	}
	if len(cd) > 0 {
		days = append(days, cd.Reverse())
	}
	return days
}

func (els Els) DayOrder() Els {
	nl := Els{}
	for _, d := range els.Days() {
		for _, e := range d {
			nl = append(nl, e)
		}
	}
	return nl
}

func (els Els) Exclude() Els {
	l := Els{}
	for _, e := range els {
		i := InfoSafe(e)
		if i["hidden"] == "true" || i["exclude"] == "true" {
			continue
		}
		l = append(l, e)
	}
	return l
}

func (els Els) Public() Els {
	l := Els{}
	for _, e := range els {
		i := InfoSafe(e)
		if i["private"] == "true" {
			continue
		}
		if Type(e) == "set" {
			s := e.(*Set)
			l = append(l, &Set{
				File:  s.File,
				Date:  s.Date,
				Info:  s.Info,
				Cover: s.Cover,
				Els:   s.Els.Public(),
			})
			continue
		}
		l = append(l, e)
	}
	return l
}

func (els Els) Reverse() Els {
	n := Els{}
	for i := len(els) - 1; i >= 0; i-- {
		n = append(n, els[i])
	}
	return n
}

func isNewDay(current, before interface{}) bool {
	cdate, err := DateSafe(current)
	if err != nil {
		return true
	}
	bdate, err := DateSafe(before)
	if err != nil {
		return true
	}
	if cdate.Hour() <= 5 {
		cdate = cdate.Add(-(time.Hour * 5))
	}
	if bdate.Hour() <= 5 {
		bdate = bdate.Add(-(time.Hour * 5))
	}
	if bdate.Day() != cdate.Day() {
		return true
	}
	return false
}


