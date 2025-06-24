package text

import (
	"g.rg-s.com/sera/go/entry/tools/blur"
	"g.rg-s.com/sera/go/entry/tools/script"
)

func (t *Text) Blur() *Text {
	t = t.Copy()

	langs := script.LangMap{}

	for k := range t.Script.LangMap {
		langs[k] = blur.ReplaceText(t.raw[k], k)
	}

	notes := langs.RenderAndExtract()

	notes.ApplyMarkdown()
	langs.ApplyMarkdown()

	BlurHyphenate(langs)

	s := &script.Script{
		LangMap:   langs,
		Footnotes: notes,
	}

	t.Script = s

	return t
}

func BlurHyphenate(langs script.LangMap) script.LangMap {
	for k, v := range langs {
		langs[k] = blur.Hyphenate(v)
	}
	return langs
}
