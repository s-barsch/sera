package hyph

import (
	"io"
	"os"
	//"github.com/akavel/go-hyphen"
	//"bufio"
	//"strings"
	"bytes"
	"unicode"
	"unicode/utf8"
)

func HyphenateText(str, lang string) (string, error) {
	err := checkEngine()
	if err != nil {
		return "", err
	}

	p := langs[lang]

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
	return text.String(), nil
}

func (p *Patterns) HyphenateWord(word string) string {
	points := p.FindHyphens(word)
	if len(points) == 0 {
		return word
	}
	nw := bytes.Buffer{}
	j := 0
	for _, r := range word {
		nw.WriteRune(r)
		if points[j]%2 != 0 && len(word) > j+1 {
			nw.WriteString("&shy;")
		}
		j++
	}
	return nw.String()
}

var langs = map[string]*Patterns{}

// TODO: create an object for this.
func checkEngine() error {
	var path = "/home/stef/go/src/stferal/pkg/entry/hyph"

	if langs["en"] == nil {
		f, err := os.Open(path + "/hyph-uk.dic")
		if err != nil {
			return err
		}
		defer f.Close()
		p, err := ParsePatterns(io.Reader(f))
		if err != nil {
			return err
		}
		langs["en"] = p
	}

	if langs["de"] == nil {
		f, err := os.Open(path + "/hyph-de.dic")
		if err != nil {
			return err
		}
		defer f.Close()
		p, err := ParsePatterns(io.Reader(f))
		if err != nil {
			return err
		}
		langs["de"] = p
	}

	return nil
}
