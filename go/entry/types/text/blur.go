package text

import (
	"g.sacerb.com/sacer/go/entry/tools/blur"
	"g.sacerb.com/sacer/go/entry/tools/script"
)

func (t *Text) Blur() *Text {
	t = t.Copy()

	langs := script.LangMap{}

	for k := range t.Script.Langs {
		langs[k] = blur.ReplaceText(t.raw[k], k)
	}

	notes := langs.RenderAndExtract()

	notes.ApplyMarkdown()
	langs.ApplyMarkdown()

	BlurHyphenate(langs)

	s := &script.Script{
		Langs:     langs,
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
