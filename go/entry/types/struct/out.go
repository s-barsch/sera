package stru

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

// base
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


/*
// custom

func (s *Struct) Path(lang string) string {
	path := ""
	for _, c := range s.Chain(lang) {
		path += "/" + c.Slug
	}
	return path
}

func (s *Struct) Chain(lang string) []*chain {
	c := &chain{
		Slug:  s.Name(lang),
		Title: s.Info.Title(lang),
	}
	if s.Mother == nil {
		return []*chain{c}
	}
	return append(s.Mother.Chain(lang), c)
}

*/
