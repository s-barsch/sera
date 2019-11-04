package entry

import (
	"os"
	"path/filepath"
	"sort"
)

func ReadSets(path string) (Sets, error) {

	sets := Sets{}

	scanSets := func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fi.IsDir() {
			return nil
		}

		if isForbiddenDir(p) {
			return filepath.SkipDir
		}

		if isDontIndex(p) {
			return nil
		}

		s, err := NewSet(p, nil)
		if err != nil {
			return err
		}

		sets = append(sets, s)

		return nil
	}

	err := filepath.Walk(path, scanSets)
	if err != nil {
		return nil, err
	}

	sort.Sort(SetDesc(sets))

	return sets, nil
}
