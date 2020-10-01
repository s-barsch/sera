package text

func (t *Text) Text(lang string) string {
	return t.Script.Langs[lang]
}
