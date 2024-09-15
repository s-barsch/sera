package script

import (
	"bytes"
	"fmt"
	"unicode/utf8"

	"g.rg-s.com/sacer/go/entry/tools"
	"g.rg-s.com/sacer/go/entry/tools/markup"
)

type Script struct {
	Langs     LangMap
	Footnotes Footnotes
}

func EmptyScript() *Script {
	return &Script{
		Langs:     map[string]string{},
		Footnotes: map[string][]string{},
	}
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

type LangMap map[string]string
type Footnotes map[string][]string

func RenderScript(langs LangMap) *Script {
	notes := langs.RenderAndExtract()

	langs.ApplyMarkdown()
	notes.ApplyMarkdown()

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

func (l LangMap) Copy() LangMap {
	m := map[string]string{}

	for k, v := range l {
		m[k] = v
	}

	return m
}

func (notes Footnotes) ApplyMarkdown() {
	for l := range tools.Langs {
		for i := range notes[l] {
			notes[l][i] = tools.MarkdownTrim(notes[l][i])
		}
	}
}

func (langs LangMap) ApplyMarkdown() {
	for l := range tools.Langs {
		langs[l] = tools.Markdown(langs[l])
	}
}

func (langs LangMap) RenderAndExtract() Footnotes {
	notes := map[string][]string{}

	for l := range tools.Langs {
		text, ns := markup.Render(langs[l])
		langs[l] = text
		notes[l] = ns
	}

	return notes
}

var inlineNoteTmpl = `<span class="footnotes inline-note"><span class="note"><span class="note-num">%d.</span> %v</span></span>`

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
