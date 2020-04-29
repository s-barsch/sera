package entry

import (
	"os"
	"path/filepath"
)

func NewText(path string, hold *Hold) (*Text, error) {
	if exists(enFile(path)) {
		return NewMultiText(path, hold)
	}
	return NewSingleText(path, hold)
}

func enFile(path string) string {
	return filepath.Join(filepath.Dir(path), "en", filepath.Base(path))
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
