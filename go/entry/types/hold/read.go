package hold

/*
import (
	"io/ioutil"
	p "path/filepath"
	"stferal/go/entry/helper"
)

func readEls(path string, hold *Hold) ([]interface{}, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	els := []interface{}{}
	for _, fi := range l {
		fp := p.Join(path, fi.Name())
		if fi.Name() == "cache" {
			subels, err := readEls(p.Join(fp, "1600"), hold)
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
		if helper.IsDontIndex(fp) {
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

	|*
		if strings.Contains(path, "/graph/") {
			sort.Sort(Desc(els))
		} else {
		}
	*|
	sort.Sort(Asc(els))

	return els, nil
}


func ReadCover(path string, hold *Hold) (*Image, error) {
	return NewImage(p.Join(path, "cover.jpg"), hold)
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

*/
