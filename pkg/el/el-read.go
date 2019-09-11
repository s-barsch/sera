package el

import (
	"path/filepath"
	// strings
	// "fmt"
)

/*
func ReadEls(folder string) (Els, error) {
	els := Els{}

	walkFn := func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err

		}

		if isForbiddenDir(p) {
			return filepath.SkipDir
		}

		if isDontIndex(p) {
			return nil
		}

		e, err := NewEl(p, nil)
		if err != nil {
			return err
		}
		els = append(els, e)

		if fi.IsDir() {
			return filepath.SkipDir
		}

		return nil
	}

	err := filepath.Walk(filepath.Join(data, folder), walkFn)
	if err != nil {
		return nil, err
	}

	sort.Sort(Desc(els))

	return els, nil
}
*/

func isForbiddenDir(p string) bool {
	switch filepath.Base(p) {
	case ".bot", "bot", "prv", "note", "pre", "en", "cor", ".versions", "vtt":
		return true
	case "320", "480", "1024", "1280", "1920", "2560", "dims":
		return true
	}
	return false
}

func isSysFile(fn string) bool {
	if len(fn) > 0 && fn[0] == '.' {
		return true
	}
	switch fn {
	case ".sort", ".bot", "bot", "en", "cache", "dims", "info", ".versions", "vtt":
		return true
	}
	return false
}

func isDontIndex(p string) bool {
	if isSysFile(filepath.Base(p)) {
		return true
	}
	switch filepath.Ext(p) {
	case ".log", ".tmp", ".xmp", ".info", ".bot":
		return true
	case ".jpg":
		if parent(p) != "1600" { // && !strings.Contains(p, "/index/") {
			return true
		}
	case "":
		return isHold(p)
		//return isStructureFolder(p)
	}
	return false
}

func parent(path string) string {
	return filepath.Base(filepath.Dir(path))
}

// Deprecated.
func isStructureFolder(p string) bool {
	info, err := ReadInfo(p)
	if err != nil {
		return true
	}
	if info["read"] == "false" {
		return true
	}
	return false
}
