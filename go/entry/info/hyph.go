package info

import (
	"fmt"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/hyph"
	"strings"
)

type Info map[string]string

func (i Info) Hyphenate() {
	for key, value := range i {
		switch name(key) {
		case "transcript", "alt":
			continue
		case "title":
			i.HyphenateTitle(key)
		case "caption":
			i[key] = tools.MarkdownNoP(value)
			fallthrough
		default:
			i[key] = hyph.Hyphenate(value, keyLang(key))
		}
	}
}

func (i Info) HyphenateTitle(key string) {
	// determine "title" or "title-en"
	lang := keyLang(key)

	// new key: "title-hyph-en"
	keyName := fmt.Sprintf("%v-hyph", name(key)) + langSuffix(lang)
	
	// save hyphenated title under new key
	i[keyName] = hyph.Hyphenate(i[key], lang)
}

func name(key string) string {
	i := strings.Index(key, "-")
	if i <= 0 {
		return key
	}
	return key[:i]
}

func keyLang(key string) string {
	i := strings.LastIndex(key, "-")
	if i <= 0 || i+1 >= len(key) {
		return "de"
	}
	return key[i+1:]
}

func langSuffix(lang string) string {
	switch lang {
	case "de":
		return ""
	default:
		return "-" + lang
	}
}
