package hyph

import (
	"os"
	"io"
)

func LoadPattern(path string) (*Patterns, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParsePatterns(io.Reader(f))
}
