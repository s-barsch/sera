package helper

import (
	"os"
	"path/filepath"
	"strings"
)

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
		if isDir(path) {
			return "dir"
		}
	}
	return "file"
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}


