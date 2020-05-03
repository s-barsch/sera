package stru

import (
	"fmt"
	"stferal/go/entry"
)

// /index/welt/wuestenleben-36c35dcb
func (s *Struct) Perma(lang string) string {
	return fmt.Sprintf("%v-%v", s.Path(lang), s.Hash())
}

// /index/welt/wuestenleben
func (s *Struct) Path(lang string) string {
	path := ""
	for _, c := range s.Chain(lang) {
		path += "/" + c.Slug
	}
	return path
}

type chain struct {
	Slug, Title string
}

func (s *Struct) Chain(lang string) []*chain {
	c := &chain{
		Slug:  s.Slug(lang),
		Title: s.Title(lang),
	}
	parent := typeCheck(s.Parent())
	if parent == nil {
		return []*chain{c}
	}
	return append(parent.Chain(lang), c)
}

func typeCheck(parentEntry entry.Entry) *Struct {
	parent, ok := parentEntry.(*Struct)
	if !ok {
		return nil
	}
	return parent
}

func (s *Struct) Depth() int {
	parent := typeCheck(s.Parent())
	if parent == nil {
		return 0
	}
	return 1 + parent.Depth()
}


