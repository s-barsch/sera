package el

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/bytefmt"
)

func (f *File) Base() string {
	return filepath.Base(f.Path)
}

func (f *File) BaseNoExt() string {
	return stripExt(filepath.Base(f.Path))
}

func (f *File) Ext() string {
	return filepath.Ext(f.Path)
}

func (f *File) Type() string {
	return FileType(f.Path)
}

func (f *File) Size() (string, error) {
	fi, err := os.Stat(f.Path)
	if err != nil {
		return "", err
	}
	size := bytefmt.ByteSize(uint64(fi.Size()))
	if len(size) < 1 {
		return "", fmt.Errorf("Size(): invalid size")
	}
	if x := size[len(size)-1:]; x != "B" && x != "0" {
		return size + "B", nil
	}
	return size, nil
}

func rel(root, path string) string {
	if len(path) > len(root) {
		return path[len(root):]
	}
	return path
}

func (f *File) Location(lang string) string { // (string, error) {
	return fmt.Sprintf("%v/files/%v", f.Hold.Permalink(lang), f.Base())
}

func locationPath(page, parentAcronym, filename string) string {
	// /{page}/{acronym}/files/{filename}
	return fmt.Sprintf("/%v/%v/files/%v", page, parentAcronym, filename)
}

func (f *File) Rel() string {
	i := strings.Index(f.Path, "data")
	if i <= 0 || len(f.Path) < i + 4 {
		return f.Path
	}
	return f.Path[i+4:]
}

func (f *File) Section() string {
	return strings.Split(f.Rel()[1:], "/")[0]
}
