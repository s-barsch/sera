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
	for _, slug := range s.Chain(lang) {
		path += "/" + slug
	}
	return path
}

func (s *Struct) Chain(lang string) []string {
	slug := s.Slug(lang)

	parent := typeCheck(s.Parent())
	if parent == nil {
		return []string{slug}
	}

	return append(parent.Chain(lang), slug)
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


