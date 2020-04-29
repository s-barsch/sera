package hyph

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

type Tree struct {
	Map    map[rune]*Tree
	Points []int
}

type Patterns struct {
	Exceptions map[string][]int
	Tree
}

const (
	_Nothing = iota
	_Patterns
	_Exceptions
)

// Parse TeX hyphenation patterns file, as published on http://tug.org/tex-hyphen/
func ParsePatterns(r io.Reader) (*Patterns, error) {
	b := bufio.NewReader(r)
	h := &Patterns{
		Tree:       Tree{Map: make(map[rune]*Tree)},
		Exceptions: make(map[string][]int),
	}
	state := _Nothing
	for {
		line, prefix, err := b.ReadLine()
		if state == _Nothing && err == io.EOF {
			return h, nil
		}
		if err != nil {
			return nil, err
		}
		if prefix {
			return nil, errors.New("Line too long")
		}

		if comment := bytes.IndexByte(line, '%'); comment != -1 {
			line = line[:comment]
		}
		line = bytes.Trim(line, " \t\n\r")
		if len(line) == 0 {
			continue
		}

		switch string(line) {
		case `\patterns{`:
			state = _Patterns
			continue
		case `\hyphenation{`:
			state = _Exceptions
			continue
		case `}`:
			state = _Nothing
			continue
		}

		switch state {
		case _Patterns:
			// Convert the a pattern like 'a1bc3d4' into a string of chars 'abcd'
			// and a list of points [0, 1, 0, 3, 4].
			//
			// Insert the pattern into the tree.  Each character finds a dict
			// another level down in the tree, and leaf nodes have the list of
			// points.
			t := &h.Tree
			points := []int{}
			p := 0
			for _, c := range string(line) {
				if '0' <= c && c <= '9' {
					p = int(c - '0') // TODO: can these be multidigit? if yes, oops
					continue
				}
				points = append(points, p)
				p = 0

				_, ok := t.Map[c]
				if !ok {
					t.Map[c] = &Tree{Map: make(map[rune]*Tree)}
				}
				t = t.Map[c]
			}
			points = append(points, p)
			t.Points = points
			/*
				chars = re.sub('[0-9]', '', pattern)
				points = [ int(d or 0) for d in re.split("[.a-z]", pattern) ]
				# Insert the pattern into the tree.  Each character finds a dict
				# another level down in the tree, and leaf nodes have the list of
				# points.
				t = self.tree
				for c in chars:
					if c not in t:
						t[c] = {}
					t = t[c]
				t[None] = points
			*/
		case _Exceptions:
			points := make([]int, 1, len(line)+2)
			word := make([]byte, 0, len(line))
			for i := 0; i < len(line); i++ {
				if line[i] != '-' {
					points = append(points, 0)
				} else {
					i++
					points = append(points, 1)
				}
				word = append(word, line[i])
			}
			points = append(points, 0)
			h.Exceptions[string(word)] = points
			//fmt.Println(string(word), points) // DEBUG
		}
	}
	panic("not reached")
}
