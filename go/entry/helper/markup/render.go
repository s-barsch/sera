package markup

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

var markupNames = map[string]string{
	"#":  "sc",
	"+":  "sperr",
	"%":  "strike",
	"~":  "italic",
}

var markupTags = map[string]string {
	"italic": "em",
	"strike": "span",
	"sc": "mark",
	"sperr": "em",
}

type Renderer struct {
	Notes []string
}

func (r *Renderer) renderSnippet(s string) string {

	buf := bytes.Buffer{}

	skip := false

	for len(s) > 0 {
		c, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		if skip {
			buf.WriteRune(c)
			skip = false
			continue
		}

		switch c {
		case '\\':
			skip = true
			continue
		case '/':
			if isNextChar(s, '/') {
				return buf.String()
			}
			if isNextChar(s, '*') {
				pos := closingPos(s[1:], "*/")
				if pos > 0 {
					s = s[pos+1:]
					continue
				}
			}
		case '{':
			x := r.renderNote(&buf, s)
			if x != 0 {
				s = s[x:]
				continue
			}
		case '#', '~', '%', '+':
			pos := r.renderTag(&buf, s, string(c))
			if pos != 0 {
				s = s[pos:]
				continue
			}
		}

		buf.WriteRune(c)
	}

	return buf.String()
}

func isNextChar(s string, c byte) bool {
	return len(s) >= 1 && s[0] == c
}

func class(c string) string {
	return ""
}

func (r *Renderer) renderTag(b *bytes.Buffer, str, c string) int {
	if x := closingPos(str, c); x != -1 {
		name := markupNames[c]
		tag := markupTags[name]
		b.WriteString(
			fmt.Sprintf(
				"<%v class=\"%v\">%v</%v>", 
				tag,
				name,
				r.renderSnippet(str[:x-1]),
				tag,
			),
		)
		return x
	}
	return 0
}

var noteTmpl = `<sup class="note-ref"><a href="#">{{.}}</a></sup>`

func (r *Renderer) renderNote(b *bytes.Buffer, s string) int {
	if x := closingPos(s, "}"); x != -1 {
		fmt.Sprintf(noteTmpl, len(r.Notes)+1)
		r.Notes = append(r.Notes, r.renderSnippet(s[:x-1]))
		return x
	}
	return 0
}

func closingPos(str string, closing string) int {
	char, _ := utf8.DecodeRuneInString(closing)
	closingRunes := utf8.RuneCountInString(closing)

	skip := false

	for i, r := range str {
		if skip {
			skip = false
			continue
		}

		if r == '\\' {
			skip = true
			continue
		}

		if r == char {
			if closingRunes == 1 {
				return i + 1
			}
			if len(str[i:]) < len(closing) {
				return -1
			}
			if str[i:i+len(closing)] == closing {
				return i + len(closing)
			}
		}
	}
	return -1
}

