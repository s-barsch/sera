package tree

import (
	"fmt"
	p "path/filepath"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

func readGraphInfo(path string, parent *Tree) (info.Info, error) {
	date, err := parseGraphDate(path, parent)
	if err != nil {
		return nil, err
	}

	// if not present, we use the empty object.
	i, _ := info.ReadDirInfo(path)

	i["date"] = date.Format(helper.Timestamp)

	if parent == nil {
		return i, nil
	}

	switch parent.Level() {
	case 0:
		setBothLang(i, "title", date.Format("2006"))
		setBothLang(i, "label", date.Format("06"))
	case 1:
		monthDe := helper.GermanMonths[date.Month()] // Januar
		monthEn := date.Format("January")
		i["title"] = monthDe
		i["title-en"] = monthEn
		i["label"] = helper.Abbr(monthDe)
		i["label-en"] = helper.Abbr(monthEn)
		setBothLang(i, "slug", date.Format("01"))
	}
	return i, nil
}

func setBothLang(i info.Info, key, value string) {
	i[key] = value
	i[key+"-en"] = value
}

func parseGraphDate(path string, parent *Tree) (time.Time, error) {
	if parent == nil {
		return time.Parse("2006_01_02", "1991_01_02")
	}
	dirName := p.Base(path)
	switch parent.Level() {
	case 0:
		return time.Parse("06", dirName)
	case 1:
		t, err := time.Parse("06-01", dirName)
		if err != nil {
			return t, err
		}
		// Workaround so 2005 and 2005-01 won’t collide.
		if t.Month() == 1 {
			t = t.Add(time.Second)
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("Could not determine graph tree date. %v", path)
}
