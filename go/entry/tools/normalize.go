package tools

import (
	bf "gopkg.in/russross/blackfriday.v2"
	"github.com/kennygrant/sanitize"
	"regexp"
	"strings"
)

var validSlug = regexp.MustCompile("[^a-z0-9-]+")

func Normalize(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Replace(name, "⹀", "-", -1)
	name = strings.Replace(name, " ", "-", -1)
	name = strings.ToLower(name)
	name = sanitize.Accents(name)

	name = validSlug.ReplaceAllString(name, "")
	return name
}

var bfExtensions = bf.WithExtensions(bf.HardLineBreak|bf.Footnotes|bf.DefinitionLists|bf.Strikethrough)

func Markdown(text string) string {
	return string(bf.Run([]byte(text), bf.WithNoExtensions(), bfExtensions))
}

func MarkdownNoP(text string) string {
	return ShaveParagraph(Markdown(text))
}

// shave off <p> and </p> at beginning and end
func ShaveParagraph(text string) string {
	if l := len(text); l > 8 {
		return text[3:l-5]
	}
	return text
}
