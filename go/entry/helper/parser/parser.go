package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode/utf8"
)

var MarkupChars = map[string]string{
	"*":  "em",
	"#":  "sc",
	"]":  "note",
	"@":  "mono",
	"##": "poem",
	/*
	"%":  "strike",
	"~":  "italic",
	"_":  "under",
	"+":  "sperr",
	*/
}

var MarkupTags = map[string]string {
	"em": "em",
	"italic": "em",
	"strike": "span",
	"mono": "span",
	"sc": "mark",
	"under": "mark",
}

type link struct {
	Name string
	Href string
}

func (t *Text) Render(style string) (string, error) {
	w := bytes.Buffer{}

	tmpl := fmt.Sprintf("text-%v", style)

	err := tmpls.ExecuteTemplate(&w, tmpl, t)
	if err != nil {
		return "", err
	}

	return w.String(), nil
}


// meat

func (p *parser) parseSnippet(notes map[string]string, str string) string {

	b := bytes.Buffer{}

	skip := false

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if skip {
			b.WriteRune(r)
			continue
		}

		switch r {

		case '*', '#', '@', '~', '%', '+':
			x := parseTag(&b, str, string(r))
			if x != 0 {
				str = str[x:]
				continue
			}
		}

		b.WriteRune(r)
	}

	return b.String()
}

func (p *parser) parseTag(b *bytes.Buffer, str, c string) int {
	if x := closingPosition(str, c); x != -1 {
		tag := markupTag[c]
		b.WriteString(fmt.Sprintf("<%v%v>%v</%v>", tag, class, parseSnippet(str[:x-1]), tag))
		p.printTag(
			b,
			MarkupChars[c],
			p.parseSnippet(str[:x-1]),
		)
		return x
	}
	return 0
}

func (p *parser) printLink(b *bytes.Buffer, name, href string) {
	err := tmpls.ExecuteTemplate(b, "html-link", &link{name, href})
	if err != nil {
		fmt.Println(err)
	}
}

func (p *parser) printTag(b *bytes.Buffer, tmpl string, data interface{}) {
	err := tmpls.ExecuteTemplate(
		b,
		fmt.Sprintf("%v-%v", tmpl, p.medium),
		data,
	)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *parser) parseNote(b *bytes.Buffer, str string) int {
	if x := closingPosition(str, "}"); x != -1 {
		p.printTag(b, "note-ref", len(p.Notes)+1)
		p.Notes = append(p.Notes, p.parseSnippet(str[:x-1]))
		return x
	}
	return 0
}

func (p *parser) parseHtml(b *bytes.Buffer, str string, c string) int {
	if x := closingPosition(str, c); x != -1 {
		b.WriteString("<")
		b.WriteString(str[:x-1])
		b.WriteString(">")
		return x
	}
	return 0
}

func closingPosition(str string, closing string) int {
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

/*
var latexSpecialChars = map[rune]string{
	'ä':  `{\"a}`,
	'Ä':  `{\"A}`,
	'ü':  `{\"u}`,
	'Ü':  `{\"U}`,
	'ö':  `{\"o}`,
	'Ö':  `{\"O}`,
	'ß':  `{\ss}`,
	'&':  `\&`,
	'%':  `\%`,
	'#':  `\#`,
	'$':  `\$`,
	'\\': `\textbackslash{}`,
	'^':  `\textasciicircum{}`,
	'_':  `\_`,
	'{':  `\{`,
	'}':  `\}`,
	'~':  `\textasciitilde{}`,
}


func parseLinks(lines []string) ([]string, map[string]string) {
	links := map[string]string{}

	for i, l := range lines {
		k, v := extractLink(l)
		links[k] = v
		if k != "" {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	return lines, links
}

func extractLink(l string) (key, value string) {
	if len(l) < 2 || l[0] != '[' {
		return
	}
	l = l[1:]
	if x := closingPosition(l, "]:"); x > 0 {
		key := l[:x-len("]:")]
		value := strings.TrimSpace(l[x:])
		return key, value
	}
	return
}
*/
/*
func newParser(input, medium string) *parser {
	return &parser{
		medium: medium,
		Notes:  []string{},
		// Links:  links,
	}
}
*/

			/*
		case '{':
			x := p.parseNote(&b, str)
			if x != 0 {
				str = str[x:]
				continue
			}
		case '<':
			x := p.parseHtml(&b, str, ">")
			if x != 0 {
				str = str[x:]
				continue
			}
		case '\\':
			skip = true
			continue
		case '/':
			// Check for // comments.
			if len(str) >= 1 && str[0] == '/' {
				return b.String()
			}
			// Check for / comments /.
			if len(str) >= 1 && str[0] == '*' {
				x := closingPosition(str[1:], "*\/")
				if x > 0 {
					str = str[x+1:]
					continue
				}
			}
		case '[':
			x := closingPosition(str, "][")
			if x > 0 {
				y := closingPosition(str[x:], "]")
				if y > 0 {
					end := x + y
					name := str[:x-len("][")]
					href := p.Links[str[x:end-1]]
					if href != "" {
						p.printLink(&b, name, href)
						str = str[end:]
						continue
					}
				}
			}
			x = closingPosition(str, "]")
			if x > 0 {
				name := str[:x-1]
				href := p.Links[name]
				if href != "" {
					if href != "" {
						p.printLink(&b, name, href)
						str = str[x:]
						continue
					}
				}
			}
			*/
