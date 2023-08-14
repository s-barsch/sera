package text

import (
	"bytes"
	"fmt"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/markup"
	"unicode/utf8"
)

type Script struct {
	Langs     Langs
	Footnotes Footnotes
}

func (s *Script) Is() bool {
	for _, str := range s.Langs {
		if str != "" {
			return true
		}
	}
	for _, strs := range s.Footnotes {
		if len(strs) > 0 {
			return true
		}
	}
	return false
}

type Langs map[string]string
type Footnotes map[string][]string

func RenderScript(langs Langs) *Script {
	notes := langs.OwnRender()

	langs.Markdown()
	notes.MarkdownHyphenate()

	return &Script{
		Langs:     langs,
		Footnotes: notes,
	}
}

func (s Script) Copy() *Script {
	return &Script{
		Langs:     s.Langs.Copy(),
		Footnotes: s.Footnotes.Copy(),
	}
}

func (n Footnotes) Copy() Footnotes {
	m := map[string][]string{}

	for k, v := range n {
		s := make([]string, len(v))
		copy(v, s)
		m[k] = s
	}

	return m
}

func (l Langs) Copy() Langs {
	m := map[string]string{}

	for k, v := range l {
		m[k] = v
	}

	return m
}

func (notes Footnotes) MarkdownHyphenate() {
	for l := range tools.Langs {
		for i := range notes[l] {
			notes[l][i] = tools.MarkdownNoP(notes[l][i])
		}
	}
}

func (langs Langs) Markdown() {
	for l := range tools.Langs {
		langs[l] = tools.Markdown(langs[l])
	}
}

func (langs Langs) OwnRender() Footnotes {
	notes := map[string][]string{}

	for l := range tools.Langs {
		text, ns := markup.Render(langs[l])
		langs[l] = text
		notes[l] = ns
	}

	return notes
}

var inlineNoteTmpl = `<span class="footnotes inline-note"><span class="note"><span class="note-num">%d.</span> %v</span>`

func (s *Script) NumberFootnotes(init int) {
	for lang := range tools.Langs {
		count := init
		buf := bytes.Buffer{}

		i := 0
		t := s.Langs[lang]
		for len(t) > 0 {
			c, size := utf8.DecodeRuneInString(t)
			t = t[size:]

			if c == '‡' {
				buf.WriteString(fmt.Sprintf("<span class=\"ref\">%d</span>", count))
				buf.WriteString(fmt.Sprintf(inlineNoteTmpl, i+1, s.Footnotes[lang][i]))
				i++
				count++
				continue
			}
			buf.WriteString(string(c))
		}

		s.Langs[lang] = buf.String()
	}
}
