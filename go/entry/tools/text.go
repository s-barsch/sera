package tools

import (
	//bf "gopkg.in/russross/blackfriday.v2"
	"bytes"
	"regexp"
	"strings"

	"g.rg-s.com/sacer/go/entry/tools/markup/gmext"
	"github.com/kennygrant/sanitize"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var validSlug = regexp.MustCompile(`[^a-z0-9-]+`)

func Normalize(name string) string {
	name = strings.TrimSpace(name)
	name = strings.Replace(name, "⹀", "-", -1)
	name = strings.Replace(name, " ", "-", -1)
	name = strings.ToLower(name)
	name = sanitize.Accents(name)

	name = validSlug.ReplaceAllString(name, "")
	return name
}

//var bfExtensions = bf.WithExtensions(bf.HardLineBreak|bf.Footnotes|bf.DefinitionLists|bf.Strikethrough)

func Markdown(text string) string {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Footnote,
			extension.DefinitionList,
			gmext.Asterism,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)
	b := bytes.Buffer{}
	if err := md.Convert([]byte(text), &b); err != nil {
		panic(err)
	}
	return b.String()
	//return string(bf.Run([]byte(text), bf.WithNoExtensions(), bfExtensions))
}

func MarkdownTrim(text string) string {
	return ShaveParagraph(Markdown(text))
}

// shave off <p> and </p> at beginning and end
func ShaveParagraph(text string) string {
	if l := len(text); l > 8 {
		return text[3 : l-5]
	}
	return text
}
