package image

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	p "path/filepath"
)

type dims struct {
	width, height int
}


func dimsFile(path string) string {
	return p.Join(cacheFolder(path), "dims", p.Base(path)+".txt")
}

func cacheFolder(path string) string {
	return p.Join(p.Dir(path), "cache")
}

func loadDims(path string) (*dims, error) {
	// /dir/file.jpg/cache/dims/file.jpg.txt
	b, err := ioutil.ReadFile(dimsFile(path))
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(b))
	x := strings.Index(s, "x")
	if x == -1 || len(s) < x+1 {
		return nil, fmt.Errorf("invalid dimensions %v", path)
	}

	w := s[:x]
	h := s[x+1:]

	width, err := strconv.Atoi(w)
	if err != nil {
		return nil, err
	}

	height, err := strconv.Atoi(h)
	if err != nil {
		return nil, err
	}

	return &dims{width: width, height: height}, nil
}


