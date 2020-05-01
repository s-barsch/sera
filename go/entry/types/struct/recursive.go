package entry

import (
	"os"
	"io/ioutil"
	p "path/filepath"
	"stferal/go/entry/helper"
	"stferal/go/entry/info"
	"strings"
)

func readStructs(path string, parent *Struct) ([]*Struct, error) {
	dirs, err := getStructDirs(path)
	if err != nil {
		return nil, err
	}

	structs, err := readStructDirs(dirs, parent)
	if err != nil {
		return nil, err
	}

	// TODO: sort structs

	return structs, nil
}

func isSymLink(fi os.FileInfo) bool { 
	return !(fi.Mode()&os.ModeSymlink != 0)
}

func getStructDirs(path string) ([]string, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dirs := []string{}

	for _, fi := range l {
		filepath := p.Join(path, fi.Name())


		if helper.IsSysFile(fi.Name()) {
			continue
		}

		if !IsHold(filepath) {
			continue
		}

		dirs = append(dirs, filepath)
	}
	return dirs, nil
}

func readStructDirs(dirs []string, parent *Struct) ([]*Struct, error) {
	structs := []*Struct{}
	for _, dirpath := range dirs {
		stru, err := ReadStruct(dirpath, parent)
		if err != nil {
			return nil, err
		}
		structs = append(structs, stru)
	}
	return structs, nil
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
