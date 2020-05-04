package tree

import (
	"io/ioutil"
	"os"
	p "path/filepath"
	//"stferal/go/entry"
	he "stferal/go/entry/helper"
	//"stferal/go/entry/parts/info"
	//"strings"
)

func readTrees(path string, parent *Tree) (Trees, error) {
	dirs, err := getTreeDirs(path)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "readTrees",
			Err:  err,
		}
	}

	trees, err := readTreeDirs(dirs, parent)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "readTrees",
			Err:  err,
		}
	}

	// TODO: sort trees

	return trees, nil
}

func readTreeDirs(dirs []string, parent *Tree) (Trees, error) {
	trees := []*Tree{}
	for _, dirpath := range dirs {
		t, err := ReadTree(dirpath, parent)
		if err != nil {
			return nil, &he.Err{
				Path: dirpath,
				Func: "readTreeDirs",
				Err:  err,
			}
		}
		trees = append(trees, t)
	}
	return trees, nil
}

func isSymLink(fi os.FileInfo) bool {
	return !(fi.Mode()&os.ModeSymlink != 0)
}

func getTreeDirs(path string) ([]string, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &he.Err{
			Path: path,
			Func: "getTreeDirs",
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
