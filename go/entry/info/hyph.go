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
			i[newKey(key, "display")] = displayTitle(key, value)
			i[key] = strings.Replace(value, "|", "", -1)
		case "caption":
			i[key] = tools.MarkdownNoP(value)
			fallthrough
		default:
			i[key] = hyph.Hyphenate(value, keyLang(key))
		}
	}
}

func displayTitle(key, value string) string {
	hyphed := hyph.Hyphenate(value, keyLang(key))

	return makeBrackets(hyphed)
}

func newKey(key, newName string) string {
	return fmt.Sprintf("%v-%v", name(key), newName) + langSuffix(keyLang(key))
}

func makeBrackets(title string) string {
	brackets := ""
	for _, el := range strings.Split(title, "|") {
		brackets += fmt.Sprintf("<span class=\"bracket\">%v</span> ", el)
	}
	return brackets
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
	switch x := key[i+1:]; x {
	case "de", "en":
		return x
	default:
		return "de"
	}
}

func langSuffix(lang string) string {
	switch lang {
	case "de":
		return ""
	default:
		return "-" + lang
	}
}
