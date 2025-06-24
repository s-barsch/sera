package tools

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Langs = map[string]string{
	"de": "Deutsch",
	"en": "English",
}

func Title(str string) string {
	return cases.Title(language.German).String(str)
}
