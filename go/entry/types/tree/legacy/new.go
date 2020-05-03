package hold

import (
	"stferal/go/entry/info"
	"stferal/go/entry/types/file"
	"time"
)

/*
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"stferal/go/entry/helper"
	"strings"
)
*/

type Tree struct {
	Parent *Tree
	File   *file.File

	Date time.Time
	Info info.Info

	Entries []*Entry
	Trees   Trees
}

type Trees []*Tree

/*
func ReadHold(path string, mother *Hold) (*Hold, error) {
	file, err := file.NewFile(path)
	if err != nil {
		return nil, err
	}

	i := info.Info{}

	if !isGraph(path) {
		info, err := info.ReadInfo(path)
		if err != nil {
			return nil, err
		}
		i = info
	} else {
		info, err := readGraphHoldInfo(path, file, mother)
		if err != nil {
			return nil, err
		}
		i = info
	}

	date, err := helper.ParseDate(i["date"])
	if err != nil {
		return nil, helper.InvalidDateErr(path, err)
	}

	h := &Hold{
		Mother: mother,
		File:   file,

		Date: date,
		Info: i,
	}

	els, err := readEls(path, h)
	if err != nil {
		return nil, err
	}

	holds, err := readHolds(path, h)
	if err != nil {
		return nil, err
	}

	h.Entries = els
	h.Holds = holds

	return h, nil
}

func readHolds(path string, mother *Hold) ([]*Hold, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	holds := []*Hold{}

	for _, f := range l {
		fp := filepath.Join(path, f.Name())
		if !(f.Mode()&os.ModeSymlink != 0) {
			if !f.IsDir() {
				continue
			}
		}
		if helper.IsSysFile(f.Name()) {
			continue
		}
		if true {
		//if !isHold(fp) {
			panic("not implemented")
			continue
		}
		h, err := ReadHold(fp, mother)
		if err != nil {
			return nil, err
		}
		// Holds that are completely empty are ommited.
		if len(h.Holds) == 0 && len(h.Entries) == 0 {
			continue
		}
		holds = append(holds, h)
	}

	sorted, err := sortHolds(path, holds)
	if err == nil {
		// ignore
		return sorted, nil
	}

	return holds, nil
}

func sortHolds(path string, holds []*Hold) ([]*Hold, error) {
	b, err := ioutil.ReadFile(path + "/.sort")
	if err != nil {
		return nil, err
	}
	l := strings.Split(strings.TrimSpace(string(b)), "\n")

	helper.ReverseStrings(l)

	for _, sortElement := range l {
		for i, h := range holds {
			if h.File.Base() == sortElement {
				cut := holds[i]
				holds = append([]*Hold{cut}, append(holds[:i], holds[i+1:]...)...)
			}
		}
	}
	return holds, nil
}

func isGraph(path string) bool {
	return strings.Contains(path, "/graph")
}

func getGraphHoldDate(path string, mother *Hold) (time.Time, error) {
	if mother == nil {
		return time.Parse("2006_01_02", "1991_01_02")
	}
	if mother.Depth() == 0 {
		return time.Parse("06", filepath.Base(path))
	}
	if mother.Depth() == 1 {
		t, err := time.Parse("06-01", filepath.Base(path))
		if err != nil {
			return t, err
		}
		// Workaround so 2005 and 2005-01 won’t collide.
		if t.Month() == 1 {
			t = t.Add(time.Second)
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("Could not determine graph hold date. %v", path)
}

func readGraphHoldInfo(path string, file *file.File, mother *Hold) (info.Info, error) {
	date, err := getGraphHoldDate(path, mother)
	if err != nil {
		return nil, err
	}

	i, err := info.ReadInfo(path)
	if err != nil {
		// ignore
	}

	i["read"] = "false"
	i["date"] = date.Format(helper.Timestamp)

	if mother != nil && mother.Depth() == 0 {
		file.Id = date.Format("2006")
		i["title"] = file.Id
		i["title-en"] = file.Id
		i["label"] = date.Format("06")
		i["label-en"] = date.Format("06")
	}

	if mother != nil && mother.Depth() == 1 {
		file.Id = date.Format("01")
		i["title"] = helper.GermanMonths[date.Month()]
		i["title-en"] = date.Format("January")
		i["slug"] = file.Id
		i["slug-en"] = file.Id
		i["label"] = helper.Abbr(i["title"])
		i["label-en"] = helper.Abbr(i["title-en"])
	}

	return i, nil
}

*/
