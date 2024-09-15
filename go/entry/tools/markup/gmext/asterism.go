package gmext

import (
	"unicode/utf8"

	"g.rg-s.com/sacer/go/entry/tools/markup/gmext/ast"
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type asterismParser struct {
}

var defaultAsterismParser = &asterismParser{}

func NewAsterismParser() parser.BlockParser {
	return defaultAsterismParser
}

func (b *asterismParser) Trigger() []byte {
	return []byte("⁂")
}

func (b *asterismParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {
	line, _ := reader.PeekLine()
	r, w := utf8.DecodeRune(line)
	if r == '⁂' {
		asterism := &ast.Asterism{}
		asterism.Offset = w
		return asterism, parser.NoChildren
	}

	return nil, parser.NoChildren
}

func (b *asterismParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	return parser.Close
}

func (b *asterismParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *asterismParser) CanInterruptParagraph() bool {
	return true
}

func (b *asterismParser) CanAcceptIndentedLine() bool {
	return false
}

type AsterismHTMLRenderer struct {
	html.Config
}

func NewAsterismHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &AsterismHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *AsterismHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindAsterism, r.renderAsterism)
}

// DefinitionListAttributeFilter defines attribute names which dl elements can have.
var DefinitionListAttributeFilter = html.GlobalAttributeFilter

func (r *AsterismHTMLRenderer) renderAsterism(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString(`<div class="asterism">*</div>`) // ⁂
	}
	return gast.WalkContinue, nil
}

type asterism struct {
}

var Asterism = &asterism{}

func (e *asterism) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(NewAsterismParser(), 101),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewAsterismHTMLRenderer(), 500),
	))
}
