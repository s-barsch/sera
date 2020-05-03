package stru

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

// base
func (e *Struct) Parent() entry.Entry {
	return e.parent
}

func (e *Struct) File() *file.File {
	return e.file
}

func (e *Struct) Id() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Struct) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Struct) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Struct) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Struct) Date() time.Time {
	return e.date
}

func (e *Struct) Info() info.Info {
	return e.info
}


// custom

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
		Slug:  s.Name(lang),
		Title: s.Title(lang),
	}
	parent := typeCheck(s.Parent)
	if parent == nil {
		return []*chain{c}
	}
	return append(parent.Chain(lang), c)
}

func (s *Struct) Depth() int {
	parent := typeCheck(s.Parent)
	if parent == nil {
		return 0
	}
	return 1 + parent.Depth()
}

func (s *Struct) Name(lang string) string {
	if slug := s.Info().Slug(lang); slug != "" {
		return slug
	}
	return s.Id()
}

func typeCheck(e entry.Entry) *Struct {
	parent, ok := e.Parent.(*Struct)
	if !ok {
		nil
	}
	return parent
}
