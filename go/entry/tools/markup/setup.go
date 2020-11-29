package markup

import (
	//"bufio"
	//"strings"
	//"bytes"
)

func Render(input string) (string, []string) {
	//s := bufio.NewScanner(strings.NewReader(input))

	r := &Renderer{
		Footnotes: []string{},
	}

	/*
	buf := bytes.Buffer{}

	for s.Scan() {
		buf.WriteString(r.renderSnippet(s.Text()))
		buf.WriteString("\n")
	}
	*/

	return r.renderSnippet(input), r.Footnotes
}

/*
func splitInput(input string) []string {
	s := bufio.NewScanner(strings.NewReader(input))

	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
*/

/*
func (p *Parser) Parse(input string) *Text {
	ls := splitInput(input)

	lines := []string{}
	for _, l := range ls {
		tt.WriteString(p.parseSnippet(l))
	}

	*/
	/*
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

	t.Footnotes = p.Footnotes
	*/

	/*
	return &Text{
		Footnotes: p.Footnotes,
		Lines: lines,
	}
}
*/

/*
func NewParser(path string, medium) (*Parser, error) {
	fm := template.FuncMap{
		"plus1": func(x int) int {
			return x + 1
		},
	}
	t, err := template.New("").Funcs(fm).ParseGlob("")
	if err != nil {
		return nil, err
	}

	return &Parser{
		Template: t,
		Footnotes: []string{},
	}, nil
}
*/


