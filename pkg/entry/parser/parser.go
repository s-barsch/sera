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
	"~":  "italic",
	"_":  "under",
	"+":  "sperr",
	"@":  "mono",
	"##": "poem",
	"%":  "strike",
}

// medium: html or tex
// style:  lines or indent

type Text struct {
	Lines []string
	Notes []string
}

type parser struct {
	Input []string
	Notes []string
	Links map[string]string

	medium string
}

// strip comments
// render markup && special chars
// hyphenate

var tmpls *template.Template

func MarkupText(input string, style string) (string, error) {
	err := checkTemplate()
	if err != nil {
		return "", err
	}

	text, err := ParseText(input, "html", style)
	if err != nil {
		return "", err
	}

	return RenderText(text, style)
}

func RenderText(t *Text, style string) (string, error) {
	w := bytes.Buffer{}

	tmpl := fmt.Sprintf("text-%v", style)

	err := tmpls.ExecuteTemplate(&w, tmpl, t)
	if err != nil {
		return "", err
	}

	return w.String(), nil
}

func newParser(input, medium string) *parser {
	s := bufio.NewScanner(strings.NewReader(input))

	lines := []string{}

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	clean, links := parseLinks(lines)

	lines = clean

	lines = hyphenateLines(lines, medium)

	return &parser{
		medium: medium,
		Input:  lines,
		Notes:  []string{},
		Links:  links,
	}
}

func hyphenateLines(lines []string, medium string) []string {
	for i, l := range lines {
		buf := bytes.Buffer{}
		s := bufio.NewScanner(strings.NewReader(l))
		s.Split(bufio.ScanWords)
		j := 0
		for s.Scan() {
			if j > 0 {
				buf.WriteString(" ")
			}
			buf.WriteString(hyphenateWord(s.Text(), medium))
			j++
		}
		lines[i] = buf.String()
	}
	return lines
}

func hyphenateWord(word string, medium string) string {
	return word
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
	if x := closingStr(l, "]:"); x > 0 {
		key := l[:x-len("]:")]
		value := strings.TrimSpace(l[x:])
		return key, value
	}
	return
}

func ParseText(input, medium, style string) (*Text, error) {

	t := &Text{Lines: []string{}}

	p := newParser(input, medium)

	for _, l := range p.Input {
		t.Lines = append(t.Lines, p.parseSnippet(l))
	}

	for i := 0; i < len(t.Lines); {
		if i > 0 && t.Lines[i] != "" && t.Lines[i-1] != "" {
			t.Lines[i-1] += "<br>" + t.Lines[i]
			t.Lines = append(t.Lines[:i], t.Lines[i+1:]...)
			continue
		}
		i++
	}

	for i, l := range t.Lines {
		if l == "" {
			if len(t.Lines) > i {
				t.Lines = append(t.Lines[:i], t.Lines[i+1:]...)
				continue
			}
			t.Lines = t.Lines[:i]
		}
	}

	t.Notes = p.Notes

	return t, nil

}

func (p *parser) parseSnippet(str string) string {

	b := bytes.Buffer{}

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		switch r {
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
		case '/':
			// Check for // comments.
			if len(str) >= 1 && str[0] == '/' {
				return b.String()
			}
			// Check for /* comments */.
			if len(str) >= 1 && str[0] == '*' {
				x := closingStr(str[1:], "*/")
				if x > 0 {
					str = str[x+1:]
					continue
				}
			}
		case '[':
			x := closingStr(str, "][")
			if x > 0 {
				y := closingStr(str[x:], "]")
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
			x = closingStr(str, "]")
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
		case '*', '#', '@', '~', '%', '+':
			x := p.parseTag(&b, str, string(r))
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
	if x := closingStr(str, c); x != -1 {
		p.printTag(
			b,
			MarkupChars[c],
			p.parseSnippet(str[:x-1]),
		)
		return x
	}
	return 0
}

type link struct {
	Name string
	Href string
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
	if x := closingStr(str, "}"); x != -1 {
		p.printTag(b, "note-ref", len(p.Notes)+1)
		p.Notes = append(p.Notes, p.parseSnippet(str[:x-1]))
		return x
	}
	return 0
}

func (p *parser) parseHtml(b *bytes.Buffer, str string, c string) int {
	if x := closingStr(str, c); x != -1 {
		b.WriteString("<")
		b.WriteString(str[:x-1])
		b.WriteString(">")
		return x
	}
	return 0
}

func closingStr(str string, closing string) int {
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

// TODO: create an object for this.
func checkTemplate() error {
	if tmpls != nil {
		return nil
	}
	fm := template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
	}
	t, err := template.New("").Funcs(fm).ParseGlob("/home/stef/go/src/stferal/pkg/entry/parser/tags.html")
	if err != nil {
		return err
	}

	tmpls = t

	return nil
}

/*
	for i < l {
		r, w := utf8.DecodeRune(d[i:])
		i += w
		switch r {
		case '/':
			if i < len(d) && d[i] == '/' {
				return b.Bytes()
			}
			if len(d) == len("/\n") {
				return []byte("<br>")
			}
			if i < len(d) && d[i] == ' ' || i == len(d) { // && i > 2 && d[i-2] != '*' {
				i += 1
				// TODO: medium
				b.WriteString("<br>")
				continue
			} else if i < len(d) && d[i] == '*' {
				// comment like this: /* comment * /
				x := closing(d[i:], "/")
				if x != 0 && x > 1 && d[i+x-1] == '*' {
					i += x + 1
					continue
				}
			}
		case '\\':
			if i+2 < l {
				next, w := utf8.DecodeRune(d[i:])
				i += w
				b.WriteRune(next)
				continue
			}
		case 'ä', 'Ä', 'ü', 'Ü', 'ö', 'Ö', 'ß', '&', '$', '^':
			if p.medium != "tex" {
				break
			}
			b.WriteString(latexSpecialChars[r])
			continue
		case '|':
			b.WriteString("&shy;")
			continue
		case '{':
			x := p.parseNote(&b, d[i:])
			if x != 0 {
				i += x
				continue
			}
		case '<':
			x := p.parseHtml(&b, d[i:], ">")
			if x != 0 {
				i += x
				continue
			}
		case '[':
			x := p.parseTag(&b, d[i:], "]")
			if x != 0 {
				i += x
				continue
			}
		case '*', '#', '@', '_', '~', '+', '%':
			if i < len(d) && d[i] == '#' {
				x := p.parseTag(&b, d[i:], "##")
				if x != 0 {
					i += x
					continue
				}
			}
			x := p.parseTag(&b, d[i:], string(r))
			if x != 0 {
				i += x
				continue
			}
		}
		b.WriteRune(r)
	}
	return b.Bytes()
}
*/

/*
func (p *parser) parseTag(b *bytes.Buffer, d []byte, c string) int {
	if x := closing(d, c); x != -1 {
		p.printTag(
			b,
			MarkupChars[c],
			p.parseSnippet(d[len(c)-1:x]),
		)
		return x + len(c)
	}
	return 0
}

func (p *parser) parseHtml(b *bytes.Buffer, d []byte, c string) int {
	if x := closing(d, c); x != -1 {
		b.WriteString("<")
		b.Write(d[len(c)-1 : x])
		b.WriteString(">")
		return x + len(c)
	}
	return 0
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

var validElId = regexp.MustCompile("[0-9]{6}_[0-9]{6}")

func (p *parser) parseNote(b *bytes.Buffer, d []byte) int {
	if x := closing(d, "}"); x != -1 {
		if z := len("160101_130101"); x >= z && validElId.Match(d[:z]) {
			fmt.Println("parseNote is not implemented")
			return 0
		}
		// what happens here? hier müsste eine fußnote eingefügt werden.
		b.WriteString("<sup>")
		b.WriteString(fmt.Sprintf("%d", len(p.Notes)+1))
		b.WriteString("</sup>")
		p.Notes = append(p.Notes, p.parseSnippet(d[:x]))
		return x + 1
	}
	return 0
}

func isClosing(d []byte, c string) bool {
	if len(c) > 1 {
		return false
	}
	if closing(d, c) != -1 {
		return true
	}
	return false
}

func closing(d []byte, c string) int {
	n, _ := utf8.DecodeLastRuneInString(c)
	skip := false
	i := 0
	for len(d) > 0 {
		r, size := utf8.DecodeRune(d)
		d = d[size:]
		i += size

		if skip {
			skip = false
			continue
		}

		if r == '\\' {
			skip = true
			continue
		}

		if r == n {
			return i - 1
		}
	}
	return -1
}

func renderLine(wr io.Writer, line []byte, style string) {
	if len(line) >= 1 && line[0] == '<' {
		wr.Write(line)
		return
	}
	tmpl := fmt.Sprintf("p-%v-html", style)
	err := tmpls.ExecuteTemplate(wr, tmpl, line)
	if err != nil {
		fmt.Println(err)
	}
}

*/
