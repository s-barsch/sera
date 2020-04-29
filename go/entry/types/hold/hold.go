package entry

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Hold struct {
	Mother *Hold
	File   *File

	Date time.Time
	Info Info
	Els  Els

	Holds Holds
	//Holds  []*Hold
}

func (h *Hold) Public() *Hold {
	c := &Hold{
		Mother: h.Mother,
		File:   h.File,

		Date: h.Date,
		Info: h.Info,
	}
	holds := Holds{}
	for _, hold := range h.Holds {
		if hold.Info["private"] == "true" {
			continue
		}
		holds = append(holds, hold.Public())
	}

	c.Els = h.Els.Public()
	c.Holds = holds

	return c
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

func ReadHold(path string, mother *Hold) (*Hold, error) {
	file, err := NewFile(path, mother)
	if err != nil {
		return nil, err
	}

	info := Info{}

	if !isGraph(path) {
		i, err := ReadInfo(path)
		if err != nil {
			return nil, err
		}
		info = i
	} else {
		i, err := readGraphHoldInfo(path, file, mother)
		if err != nil {
			return nil, err
		}
		info = i
	}

	date, err := ParseDate(info["date"])
	if err != nil {
		return nil, invalidDate(path, err)
	}

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

	holds, err := readHolds(path, h)
	if err != nil {
		return nil, err
	}

	h.Els = els
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
		if isSysFile(f.Name()) {
			continue
		}
		if !isHold(fp) {
			continue
		}
		h, err := ReadHold(fp, mother)
		if err != nil {
			return nil, err
		}
		/* Holds that are completely empty are ommited. */
		if len(h.Holds) == 0 && len(h.Els) == 0 {
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

func isHold(path string) bool {
	info, err := ReadInfo(path)
	if err != nil {
		if strings.Contains(path, "/graph") {
			return true
		}
		return false
	}
	if info["inline"] == "true" {
		return false
	}
	return true
}

func sortHolds(path string, holds []*Hold) ([]*Hold, error) {
	b, err := ioutil.ReadFile(path + "/.sort")
	if err != nil {
		return nil, err
	}
	l := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, sortElement := range reverse(l) {
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

func readGraphHoldInfo(path string, file *File, mother *Hold) (Info, error) {
	date, err := getGraphHoldDate(path, mother)
	if err != nil {
		return nil, err
	}

	info, err := ReadInfo(path)
	if err != nil {
		// ignore
	}

	info["read"] = "false"
	info["date"] = date.Format(Timestamp)

	if mother != nil && mother.Depth() == 0 {
		file.Id = date.Format("2006")
		info["title"] = file.Id
		info["title-en"] = file.Id
		info["label"] = date.Format("06")
		info["label-en"] = date.Format("06")
	}

	if mother != nil && mother.Depth() == 1 {
		file.Id = date.Format("01")
		info["title"] = GermanMonths[date.Month()]
		info["title-en"] = date.Format("January")
		info["slug"] = file.Id
		info["slug-en"] = file.Id
		info["label"] = Abbr(info["title"])
		info["label-en"] = Abbr(info["title-en"])
	}

	return info, nil
}
