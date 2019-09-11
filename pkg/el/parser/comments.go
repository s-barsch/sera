package parser

import (
	"unicode/utf8"
)

func RemoveHidden(str string) string {
	clean := ""
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		switch r {
		case '/':
			if len(str) > 0 && str[0] == '/' {
				i := closingStr(str, "\n")
				if i == -1 {
					return clean
				}
				str = str[i-1:]
				continue
			}
			if len(str) > 0 && str[0] == '*' {
				i := closingStr(str, "*/")
				if i != -1 {
					str = str[i:]
					continue
				}
			}
		case '{':
			i := closingStr(str, "}")
			if i != -1 {
				str = str[i:]
				continue
			}
		}

		clean += string(r)
	}
	return clean
}

/*
func advance(str string, closing string) int {
	if len(closing) == 0 {
		return -1
	}
	i := 0
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		i += size

		if r == closing[0] {
			if len(closing) > 1 {
			}
			return i
		}
	}
	return -1
}
*/
