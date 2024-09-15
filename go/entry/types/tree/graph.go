package tree

import (
	"fmt"
	p "path/filepath"
	"time"

	"g.rg-s.com/sacer/go/entry/info"
	"g.rg-s.com/sacer/go/entry/tools"
)

func readGraphInfo(path string, parent *Tree) (info.Info, error) {
	d, d2, err := parseFolderDate(path, parent)
	if err != nil {
		return nil, fmt.Errorf("readGraphInfo: %v", err)

	}

	// if not present, we use the empty object.
	i, _ := info.ReadDirInfo(path)

	i["date"] = d.Format(tools.Timestamp)

	if parent == nil {
		return i, nil
	}

	switch parent.Level() {
	case 0:
		setBothLang(i, "title", d.Format("2006"))
		setBothLang(i, "label", d.Format("06"))
	case 1:
		m := buildMonthName(d)
		i["title"] = m["de"]
		i["title-en"] = m["en"]
		i["label"] = tools.Abbr(m["de"])
		i["label-en"] = tools.Abbr(m["en"])
		setBothLang(i, "slug", d.Format("01"))

		if !d2.IsZero() {
			m = buildMonthName(d2)
			i["title"] = fmt.Sprintf("%v – %v", i["title"], m["de"])
			i["title-en"] = fmt.Sprintf("%v – %v", i["title-en"], m["en"])
			i["label"] = fmt.Sprintf("%v–%v", i["label"], tools.Abbr(m["de"]))
			i["label-en"] = fmt.Sprintf("%v–%v", i["label-en"], tools.Abbr(m["en"]))
			setBothLang(i, "slug", fmt.Sprintf("%v-%v", d.Format("01"), d2.Format("01")))
		}

	}
	return i, nil
}

func buildMonthName(date time.Time) map[string]string {
	return map[string]string{
		"de": tools.GermanMonths[date.Month()],
		"en": date.Format("January"),
	}
}

func setBothLang(i info.Info, key, value string) {
	i[key] = value
	i[key+"-en"] = value
}

func parseFolderDate(path string, parent *Tree) (time.Time, time.Time, error) {
	if parent == nil {
		d, err := time.Parse("2006_01_02", "1991_01_02")
		return d, time.Time{}, err
	}
	dirName := p.Base(path)
	switch parent.Level() {
	case 0:
		d, err := time.Parse("06", dirName)
		return d, time.Time{}, err
	case 1:
		d2 := time.Time{}
		format := "06-01"
		l := len(format)
		if isMergeTree(path) {
			d, err := time.Parse("01", dirName[l:])
			if err != nil {
				return time.Time{}, time.Time{}, err
			}
			dirName = dirName[:l]
			d2 = d
		}
		d, err := time.Parse("06-01", dirName)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		// Workaround so 2005 and 2005-01 won’t collide.
		d = d.Add(time.Second)
		return d, d2, nil
	}
	return time.Time{}, time.Time{}, fmt.Errorf("could not determine graph tree date: %v", path)
}
