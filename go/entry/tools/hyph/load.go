package hyph

import (
	"os"
	"io"
)

var lp LangPatterns

type LangPatterns map[string]*Patterns

func SetLangPatterns(langp LangPatterns) {
	lp = langp
}

func LoadPattern(path string) (*Patterns, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParsePatterns(io.Reader(f))
}
