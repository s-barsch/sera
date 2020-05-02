import (
	"fmt"
	"io/ioutil"
	"strconv"
)

type dims struct {
	width, height int
}


func dimsFile(path string) string {
	return filepath.Join(cacheFolder(path), "dims", filepath.Base(path)+".txt")
}

func cacheFolder(path string) string {
	return filepath.Join(filepath.Dir(path), "cache")
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


