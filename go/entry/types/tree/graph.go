package tree

import (
	"fmt"
	p "path/filepath"
	"sacer/go/entry/info"
	"sacer/go/entry/tools"
	"time"
)

func readGraphInfo(path string, parent *Tree) (info.Info, error) {
	date, err := parseFolderDate(path, parent)
	if err != nil {
		return nil, fmt.Errorf("readGraphInfo: %v", err)

	}

	// if not present, we use the empty object.
	i, _ := info.ReadDirInfo(path)

	i["date"] = date.Format(tools.Timestamp)

	if parent == nil {
		return i, nil
	}

	switch parent.Level() {
	case 0:
		setBothLang(i, "title", date.Format("2006"))
		setBothLang(i, "label", date.Format("06"))
	case 1:
		m := buildMonthName(date)
		i["title"] = m["de"]
		i["title-en"] = m["en"]
		i["label"] = tools.Abbr(m["de"])
		i["label-en"] = tools.Abbr(m["en"])
		setBothLang(i, "slug", date.Format("01"))

		if isMergeTree(path) {
			t, err := time.Parse("06-01", "20-12")
			if err != nil {
				fmt.Println(err)
			}
			m = buildMonthName(t)
			i["title"] = fmt.Sprintf("%v – %v", i["title"], m["de"])
			i["title-en"] = fmt.Sprintf("%v – %v", i["title-en"], m["en"])
			i["label"] = fmt.Sprintf("%v–%v", i["label"], tools.Abbr(m["de"]))
			i["label-en"] = fmt.Sprintf("%v–%v", i["label-en"], tools.Abbr(m["en"]))
			setBothLang(i, "slug", "11-12")
			fmt.Println(i)
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

func parseFolderDate(path string, parent *Tree) (time.Time, error) {
	if parent == nil {
		return time.Parse("2006_01_02", "1991_01_02")
	}
	dirName := p.Base(path)
	switch parent.Level() {
	case 0:
		return time.Parse("06", dirName)
	case 1:
		format := "06-01"
		if len(dirName) > len(format) {
			dirName = dirName[:len(format)]
		}
		t, err := time.Parse("06-01", dirName)
		if err != nil {
			return t, err
		}
		// Workaround so 2005 and 2005-01 won’t collide.
		t = t.Add(time.Second)
		return t, nil
	}
	return time.Time{}, fmt.Errorf("Could not determine graph tree date. %v", path)
}
