package info

import (
	"fmt"
	"strings"

	"g.rg-s.com/sera/go/entry/tools"
)

type Info map[string]string

func (i Info) Hyphenate() {
	for key, value := range i {
		switch name(key) {
		case "transcript", "summary", "summary-private", "alt", "slug":
			continue
		case "title":
			i[newKey(key, "display")] = makeBrackets(value)
			i[key] = strings.Replace(value, "|", "", -1)
			i[newKey(key, "hyph")] = i[key]
		case "caption":
			i[key] = tools.MarkdownTrim(value)
		}
	}
}

func newKey(key, newName string) string {
	return fmt.Sprintf("%v-%v", name(key), newName) + langSuffix(keyLang(key))
}

func makeBrackets(title string) string {
	brackets := ""
	for _, el := range strings.Split(title, "|") {
		brackets += fmt.Sprintf("<span class=\"bracket\">%v</span> ", el)
	}
	return strings.TrimSpace(brackets)
}

func name(key string) string {
	i := strings.LastIndex(key, "-")
	if i <= 0 || key[i+1:] != "en" {
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
