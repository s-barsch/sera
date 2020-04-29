package entry

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Id   string
	Path string

	ModTime time.Time
	Hold    *Hold
}

func getModTime(path string) (time.Time, error) {

	switch FileType(path) {
	case "image", "video":
		return getModTimeMedia(path)
	case "dir":
		return getModTimeDir(path, true)
	default:
		return getModTimeFile(path)
	}
}

func getModTimeFile(path string) (time.Time, error) {
	t := time.Time{}

	fi, err := os.Stat(path)
	if err != nil {
		return t, err
	}

	return fi.ModTime(), nil
}

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

func NewFile(path string, mother *Hold) (*File, error) {
	t, err := getModTime(path)
	if err != nil {
		return nil, err
	}

	id := makeId(filepath.Base(path))
	return &File{
		Id:   id,
		Path: path,

		ModTime: t,
		Hold:    mother,
	}, nil
}

func makeId(p string) string {
	return Shorten(filepath.Base(p))
}

func Shorten(n string) string {
	if len(n) > 13 {
		return n[:13]
	}
	return n
}

func loadFileInfo(path string) (Info, error) {
	path += ".info"
	_, err := os.Stat(path)
	if err == nil {
		return parseInfoFile(path)
	}
	return map[string]string{}, nil
}

func getFilenameDate(path string) (time.Time, error) {
	return time.Parse(Timestamp, Shorten(filepath.Base(path)))
}
