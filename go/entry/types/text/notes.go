package text

import (
	"sacer/go/entry/tools"
	"bytes"
	"unicode/utf8"
	"strconv"
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
	s := t.TextLangs[lang]
	for len(s) > 0 {
		c, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		if c == '‡' {
			buf.WriteString(strconv.Itoa(count))
			buf.WriteString(fmt.Sprintf("{%v}", t.Notes[lang][i]))
			i++
			count++
			continue
		}
		buf.WriteString(string(c))
	}

	t.TextLangs[lang] = buf.String()
}
