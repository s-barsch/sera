package text

func (t *Text) Text(lang string) string {
	return t.Langs[lang]
}
