package helper

import (
	"os"
	p "path/filepath"
	"sort"
	"strings"
)

func ReverseStrings(slice []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
}

func Shorten(n string) string {
	if len(n) > 13 {
		return n[:13]
	}
	return n
}

func StripExt(base string) string {
	i := strings.LastIndex(base, ".")
	if i <= 0 {
		return base
	}
	return base[:i]
}

func FileType(path string) string {
	switch p.Ext(path) {
	case ".txt", ".ltxt", ".ptxt", ".itxt":
		return "text"
	case ".mp3", ".wav":
		return "audio"
	case ".mp4":
		return "video"
	case ".jpg", ".png", ".svg":
		return "image"
	case ".html":
		return "html"
	case "":
		if IsDir(path) {
			return "dir"
		}
	}
	return "file"
}

func ParentDir(path string) string {
	return p.Base(p.Dir(path))
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}

func IsNameSys(name string) bool {
	if startsWithDot(name) {
		return true
	}

	switch name {
	case ".sort", "bot", "cache", "info", "vtt":
		return true
	}

	return false
}

func startsWithDot(name string) bool {
	if len(name) > 0 && name[0] == '.' {
		return true
	}
	return false
}

func IsDontIndex(path string) bool {
	if IsNameSys(p.Base(path)) {
		return true
	}
	switch p.Ext(path) {
	case ".log", ".tmp", ".xmp", ".info", ".bot":
		return true
	case ".jpg":
		// TODO: maybe rewrite that?
		if ParentDir(path) != "1600" {
			return true
		}
		// Let the packages determine if regular dirs should be index or not.
		// case "":
	}
	return false
}
