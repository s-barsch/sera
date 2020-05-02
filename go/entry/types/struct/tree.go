package stru

import (
	"io/ioutil"
	"os"
	p "path/filepath"
	he "stferal/go/entry/helper"
	//"stferal/go/entry/parts/info"
	//"strings"
)

func readStructs(path string, parent *Struct) ([]*Struct, error) {
	dirs, err := getStructDirs(path)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "readStructs",
			Err:  err,
		}
	}

	structs, err := readStructDirs(dirs, parent)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "readStructs",
			Err:  err,
		}
	}

	// TODO: sort structs

	return structs, nil
}

func readStructDirs(dirs []string, parent *Struct) ([]*Struct, error) {
	structs := []*Struct{}
	for _, dirpath := range dirs {
		stru, err := ReadStruct(dirpath, parent)
		if err != nil {
			return nil, &he.Err{
				Path: dirpath,
				Func: "readStructDirs",
				Err:  err,
			}
		}
		structs = append(structs, stru)
	}
	return structs, nil
}

func isSymLink(fi os.FileInfo) bool {
	return !(fi.Mode()&os.ModeSymlink != 0)
}

func getStructDirs(path string) ([]string, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &he.Err{
			Path: path, 
			Func: "getStructDirs",
			Err:  err,
		}
	}

	dirs := []string{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())

		if he.IsDontIndex(fi.Name()) {
			continue
		}

		if !fi.IsDir() {
			continue
		}

		// figure that out for graph etc
		/*
			if !IsHold(filepath) {
				continue
			}
		*/

		dirs = append(dirs, filepath)
	}
	return dirs, nil
}

/*
	if !isSymLink(fi) && !fi.IsDir() {
		if !fi.IsDir() {
			continue
		}
	}
*/

/*
	// Holds that are completely empty are ommited.
	if len(h.Holds) == 0 && len(h.Entries) == 0 {
		continue
	}
*/

/*
func IsHold(path string) bool {
	info, err := info.ReadInfo(path)
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
*/
