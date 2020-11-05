package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type Asterism struct {
	gast.BaseBlock
	Offset             int
	TemporaryParagraph *gast.Paragraph
}

// Dump implements Node.Dump.
func (n *Asterism) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindAsterism is a NodeKind of the Asterism node.
var KindAsterism = gast.NewNodeKind("Asterism")

// Kind implements Node.Kind.
func (n *Asterism) Kind() gast.NodeKind {
	return KindAsterism
}

// NewAsterism returns a new Asterism node.
func NewAsterism(offset int, para *gast.Paragraph) *Asterism {
	return &Asterism{
		Offset:             offset,
		TemporaryParagraph: para,
	}
}
