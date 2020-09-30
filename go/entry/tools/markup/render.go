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
		case ':':
			// If a slash follows a colon, comments are disabled for
			// that sequence to prevent commenting https://domain.
			if isNextChar(s, '/') {
				skip = true
				buf.WriteRune(c)
				continue
			}
		case '/':
			if isNextChar(s, '/') {
				pos := closingPos(s[1:], "\n")
				if pos > 0 {
					s = s[pos:]
					continue
				}
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
			if !isNextChar(s, byte(c)) {
				pos := r.renderTag(&buf, s, string(c))
				if pos != 0 {
					s = s[pos:]
					continue
				}
			}
		case '⁂':
			buf.WriteString("<span class=\"asterism\">⁂</span>")
			continue
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

//var noteTmpl = `<sup class="reference"><a href="#">%v</a></sup>`
var noteTmpl = "‡"

func (r *Renderer) renderNote(buf *bytes.Buffer, s string) int {
	if x := closingPos(s, "}"); x != -1 {
		//buf.WriteString(fmt.Sprintf(noteTmpl, len(r.Notes)+1))
		buf.WriteString(noteTmpl)
		r.Notes = append(r.Notes, r.renderSnippet(s[:x-1]))
		return x
	}
	return 0
}

func closingPos(s string, closing string) int {
	char, _ := utf8.DecodeRuneInString(closing)
	closingRunes := utf8.RuneCountInString(closing)

	skip := false

	for i, r := range s {
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
			if len(s[i:]) < len(closing) {
				return -1
			}
			if s[i:i+len(closing)] == closing {
				return i + len(closing)
			}
		}

		if r == '\n' {
			return -1
		}
	}
	return -1
}

