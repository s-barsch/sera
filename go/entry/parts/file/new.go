package file

import (
	"os"
	"time"
	"stferal/go/entry/helper"
)

type File struct {
	Path    string
	ModTime time.Time
}

func NewFile(path string) (*File, error) {
	mod, err := getModTime(path)
	if err != nil {
		return nil, &helper.Err{
			Path: path,
			Func: "NewFile",
			Err:  err,
		}
	}

	return &File{
		Path:    path,
		ModTime: mod,
	}, nil
}

func getModTime(path string) (time.Time, error) {
	return getModTimeFile(path)
}

func getModTimeFile(path string) (time.Time, error) {
	t := time.Time{}

	fi, err := os.Stat(path)
	if err != nil {
		return t, &helper.Err{
			Path: path,
			Func: "getModTimeFile",
			Err:  err,
		}
	}

	return fi.ModTime(), nil
}

/*
func getModTimeMedia(path string) (time.Time, error) {
	t, err := getModTimeFile(path)
	if err != nil {
		return t, err
	}

	tInfo, err := getModTimeFile(path + ".info")
	if err != nil {
		return t, nil
	}

	if tInfo.Unix() > t.Unix() {
		return tInfo, nil
	}

	return t, nil
}

func getModTimeDir(path string, recur bool) (time.Time, error) {
	t := time.Time{}

	dir, err := os.Stat(path)
	if err != nil {
		return t, err
	}

	t = dir.ModTime()

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return t, err
	}

	for _, fi := range fis {
		if filepath.Ext(fi.Name()) == ".info" && !recur {
			continue
		}

		tChild := fi.ModTime()
		if fi.IsDir() && recur {
			td, err := getModTimeDir(path+"/"+fi.Name(), false)
			if err != nil {
				return t, err
			}
			tChild = td
		}

		if tChild.Unix() > t.Unix() {
			t = fi.ModTime()
		}
	}

	return t, nil
}


func getFilenameDate(path string) (time.Time, error) {
	return time.Parse(helper.Timestamp, Shorten(filepath.Base(path)))
}
*/
