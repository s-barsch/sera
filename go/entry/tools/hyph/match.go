package hyph

import (
	"strings"
	"unicode/utf8"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Given a word, returns a list of pieces, broken at the possible
// hyphenation points. Note: returned pieces may reuse underlying storage
// of the input word.
func (h *Patterns) FindHyphens(word string) []int {
	// Short words aren't hyphenated.
	if utf8.RuneCountInString(word) <= 4 {
		return []int{}
	}

	//var points []int
	points := []int{}

	lower := strings.ToLower(word)

	// If the word is an exception, get the stored points.
	if exception, ok := h.Exceptions[lower]; ok {
		points = exception
	} else {
		/*
			work := make([]byte, 0, len(lower)+2)
			work = append(work, '.')
			work = append(work, lower...)
			work = append(work, '.')
		*/

		work := "." + lower + "."

		points = make([]int, len(work)+1)

		for i := range work {
			t := &h.Tree
			for _, r := range work[i:] {
				tTmp, ok := t.Map[r]
				if !ok {
					break
				}
				t = tTmp

				p := t.Points
				if p == nil {
					continue
				}

				for j := range p {
					points[i+j] = max(points[i+j], p[j])
				}

			}
		}

		// No hyphens in the first two chars or the last two.
		points[1], points[2] = 0, 0
		points[len(points)-2], points[len(points)-3] = 0, 0
	}

	points = points[2:]
	/*
		points[len(word)-1] = -1 // loop terminator
	*/

	return points

	/*
		// Examine the points to build the pieces list.
		pieces := [][]byte{}
		points = points[2:]
		points[len(word)-1] = -1 // loop terminator
		last := 0
		for i := range word {
			if points[i]%2 == 0 {
				continue
			}
			pieces = append(pieces, word[last:i+1])
			last = i + 1
		}
		return pieces
	*/
}
