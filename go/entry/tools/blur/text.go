package blur

import (
	"math/rand"
	"unicode"
	"unicode/utf8"
	"bytes"
	"time"
)

func ReplaceText(str, lang string) string {
	buf := bytes.Buffer{}
	rand.Seed(time.Now().UnixNano())
	charLen := len(chars)
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if unicode.IsLetter(r) {
			x := chars[rand.Intn(charLen)]
			if unicode.IsUpper(r) {
				x -= 32
			}
			buf.WriteRune(x)
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

// this is to prioritize vowels
var chars = []rune {
	'a',
	'a',
	'a',
	'e',
	'e',
	'e',
	'i',
	'i',
	'i',
	'o',
	'o',
	'o',
	'u',
	'u',
	'u',
	'b',
	'c',
	'd',
	'h',
	'j',
	'k',
	'l',
	'm',
	'n',
	'p',
	'r',
	's',
	't',
}

func Hyphenate(str string) string {
	buf := bytes.Buffer{}
	rand.Seed(time.Now().UnixNano())
	c := 0
	html := false
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if r == '>' {
			html = false
			buf.WriteRune(r)
			continue
		}

		if r == '<' {
			html = true
			buf.WriteRune(r)
			continue
		}

		if html {
			buf.WriteRune(r)
			continue
		}

		if unicode.IsLetter(r) {
			if c >= 3 {
				buf.WriteString("&shy;")
				c = 0
			}
			buf.WriteRune(r)
			c++
			continue
		}
		c = 0
		buf.WriteRune(r)
	}
	return buf.String()
}

