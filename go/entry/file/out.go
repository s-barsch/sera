package file

import (
	"path/filepath"
	"stferal/go/entry/helper"
)

func (f *File) Base() string {
	return filepath.Base(f.Path)
}

func (f *File) BaseNoExt() string {
	return helper.StripExt(filepath.Base(f.Path))
}

func (f *File) Ext() string {
	return filepath.Ext(f.Path)
}

func (f *File) Type() string {
	return helper.FileType(f.Path)
}

/*
	"fmt"
	"code.cloudfoundry.org/bytefmt"
	"strings"
	"os"
*/

/*
func (f *File) Section() string {
	return strings.Split(f.Rel()[1:], "/")[0]
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
*/
