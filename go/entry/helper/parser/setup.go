package parser

type Text struct {
	Lines []string
	Notes []string
}

type Parser struct {
	Template *template.Template

	Notes []string
	Links map[string]string
}

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
		Notes: []string{},
	}, nil
}

//var tmpls *template.Template

func (p *Parser) MarkUp(input string) (string, error) {
	t := p.Parse(input)

	s := bufio.NewScanner(strings.NewReader(input))

	lines := []string{}
	for s.Scan() {
		p.WriteString(s
		lines = append(lines, s.Text())
	}
	return lines

	return t.Render()
}

func splitInput(input string) []string {
	s := bufio.NewScanner(strings.NewReader(input))

	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func (p *Parser) Parse(input string) *Text {
	ls := splitInput(input)

	lines := []string{}
	for _, l := range ls {
		tt.WriteString(p.parseSnippet(l))
	}

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

	t.Notes = p.Notes
	*/

	return &Text{
		Notes: p.Notes,
		Lines: lines,
	}
}


