package hyph

import (
	//"github.com/akavel/go-hyphen"
	//"bufio"
	//"strings"
	"bytes"
	"unicode"
	"unicode/utf8"
)

func (p *Patterns) HyphenateText(str string) string {
	word := bytes.Buffer{}
	text := bytes.Buffer{}

	insideWord := false
	insideHTML := false

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if r == '<' {
			if insideWord {
				insideWord = false
				text.WriteString(p.HyphenateWord(word.String()))
				word.Reset()
			}
			insideHTML = true
			text.WriteRune(r)
			continue
		}

		if r == '>' && insideHTML {
			insideHTML = false
			text.WriteRune(r)
			continue
		}

		if insideHTML {
			text.WriteRune(r)
			continue
		}

		if unicode.IsLetter(r) {
			insideWord = true
			word.WriteRune(r)
			continue
		}

		if insideWord {
			text.WriteString(p.HyphenateWord(word.String()))
			insideWord = false
			word.Reset()
		}
		text.WriteRune(r)
	}
	text.WriteString(p.HyphenateWord(word.String()))
	return text.String()
}

func (p *Patterns) HyphenateWord(word string) string {
	points := p.FindHyphens(word)
	if len(points) == 0 {
		return word
	}
	nw := bytes.Buffer{}
	l := utf8.RuneCountInString(word)
	for i, r := range word {
		nw.WriteRune(r)
		if points[i]%2 != 0 && l > i+1 {
			nw.WriteString("&shy;")
		}
	}
	return nw.String()
}



