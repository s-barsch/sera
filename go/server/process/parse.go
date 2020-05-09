package process

func newParser(tmpl string) (*parser.Parser, error) {
	tmpl := "/home/stef/go/src/stferal/go/entry/parser/tags.html"
	return parser.NewParser(tmpl)
}

type Processer struct {
	Parser *parse.Parser
	Hypher map[string]*hyphe
	Lang   []string
}

var langs = []string{
	"de",
	"en",
}

func ParseTexts(tmpl string, entries entry.Entries) error {
	p, err := newParser(tmpl)
	if err != nil {
		return err
	}
	for _, e := range entries {
		tx, ok := e.(*text.Text)
		if !ok {
			continue
		}
	}
}



func (p *Processer) ProcessText(p *parse.Parser, tx *textText) error {
	for _, l := range langs {
		render := p.Parser.Markup(tx.Raw(l))
		tx.SetText(render, lang)
	}
}
