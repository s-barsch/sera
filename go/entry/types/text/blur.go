package text

import (
	"sacer/go/entry/tools/blur"
)

func (t *Text) Blur() *Text {
	t = t.Copy()

	langs := Langs{}

	for k, _ := range t.Script.Langs {
		langs[k] = blur.ReplaceText(t.raw[k], k)
	}

	notes := langs.OwnRender()
	notes.MarkdownHyphenate()

	langs.Markdown()
	langs.BlurHyphenate()
	
	s := &Script{
		Langs: langs,
		Notes: notes,
	}

	t.Script = s

	return t
}

func (langs Langs) BlurHyphenate() {
	for k, v := range langs {
		langs[k] = blur.Hyphenate(v)
	}
}

