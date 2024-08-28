package text

/*
import (
	"g.sacerb.com/sacer/go/entry/tools"
	"bytes"
	"unicode/utf8"
	"fmt"
)

func (t *Text) RenderFootnotes(count int) {
	for lang, _ := range tools.Langs {
		t.RenderFootnotesLang(count, lang)
	}
}

func (t *Text) RenderFootnotesLang(count int, lang string) {
	buf := bytes.Buffer{}

	i := 0
	s := t.Script.Langs[lang]
	for len(s) > 0 {
		c, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		if c == '‡' {
			buf.WriteString(fmt.Sprintf("<span class=\"ref\">%d</span>", count))
			buf.WriteString(fmt.Sprintf("<span class=\"inline-note\">%v</span>", t.Script.Notes[lang][i]))
			i++
			count++
			continue
		}
		buf.WriteString(string(c))
	}

	t.Script.LangMap[lang] = buf.String()
}
*/
