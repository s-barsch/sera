// Code generated by go generate; DO NOT EDIT.

package stru

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

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

func (e *Struct) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}
