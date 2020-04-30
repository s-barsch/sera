package helper

import (
	"os"
	"path/filepath"
	"strings"
	"sort"
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
	switch filepath.Ext(path) {
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
	return filepath.Base(filepath.Dir(path))
}

/*
func IsHold(path string) bool {
	info, err := os.Stat(path)
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

func IsForbiddenDir(p string) bool {
	switch filepath.Base(p) {
	case ".bot", "bot", "prv", "note", "pre", "en", "cor", ".versions", "vtt":
		return true
	case "320", "480", "1024", "1280", "1920", "2560", "dims":
		return true
	}
	return false
}

func IsSysFile(path string) bool {
	fn := filepath.Base(path)
	if len(fn) > 0 && fn[0] == '.' {
		return true
	}
	switch fn {
	case ".sort", ".bot", "bot", "en", "cache", "dims", "info", ".versions", "vtt":
		return true
	}
	return false
}

func IsDontIndex(path string) bool {
	if IsSysFile(path) {
		return true
	}
	switch filepath.Ext(path) {
	case ".log", ".tmp", ".xmp", ".info", ".bot":
		return true
	case ".jpg":
		if ParentDir(path) != "1600" { // && !strings.Contains(p, "/index/") {
			return true
		}
	case "":
		panic("not implemented yets")
		return true
		//return isHold(p)
		//return isStructureFolder(p)
	}
	return false
}

/*
// Deprecated.
func IsStructureFolder(p string) bool {
	info, err := ReadInfo(p)
	if err != nil {
		return true
	}
	if info["read"] == "false" {
		return true
	}
	return false
}
*/
