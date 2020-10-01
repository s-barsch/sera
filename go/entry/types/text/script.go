package text

import (
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/hyph"
	"sacer/go/entry/tools/markup"
	"bytes"
	"unicode/utf8"
	"fmt"
)

type Script struct {
	Langs Langs
	Notes Notes
}

type Langs map[string]string
type Notes map[string][]string

func RenderScript(langs Langs) (*Script) {
	notes := langs.OwnRender()

	langs.Markdown()
	langs.Hyphenate()
	notes.MarkdownHyphenate()

	return &Script{
		Langs: langs,
		Notes: notes,
	}
}


func (notes Notes) MarkdownHyphenate() {
	for l, _ := range tools.Langs {
		for i, _ := range notes[l] {
			notes[l][i] = tools.MarkdownNoP(notes[l][i])
			notes[l][i] = hyph.Hyphenate(notes[l][i], l)
		}
	}
}

func (langs Langs) Hyphenate() {
	for l, _ := range tools.Langs {
		langs[l] = hyph.Hyphenate(langs[l], l)
	}
}

func (langs Langs) Markdown() {
	for l, _ := range tools.Langs {
		langs[l] = tools.Markdown(langs[l])
	}
}

func (langs Langs) OwnRender() Notes {
	notes := map[string][]string{}

	for l, _ := range tools.Langs {
		text, ns := markup.Render(langs[l])
		langs[l] = text
		notes[l] = ns
	}

	return notes
}

func (s *Script) NumberFootnotes(count int) {
	for lang, _ := range tools.Langs {
		buf := bytes.Buffer{}

		i := 0
		t := s.Langs[lang]
		for len(t) > 0 {
			c, size := utf8.DecodeRuneInString(t)
			t = t[size:]

			if c == '‡' {
				buf.WriteString(fmt.Sprintf("<span class=\"ref\">%d</span>", count))
				buf.WriteString(fmt.Sprintf("<span class=\"inline-note\">%v</span>", s.Notes[lang][i]))
				i++
				count++
				continue
			}
			buf.WriteString(string(c))
		}

		s.Langs[lang] = buf.String()
	}
}
