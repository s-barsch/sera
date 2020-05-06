// Code generated by go generate; DO NOT EDIT.

package image

import (
	"fmt"

	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Image) Type() string {
	return "image"
}

func (e *Image) Parent() entry.Entry {
	return e.parent
}

func (e *Image) File() *file.File {
	return e.file
}

func (e *Image) Id() int64 {
	return e.date.Unix()
}

func (e *Image) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Image) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Image) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Image) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Image) Date() time.Time {
	return e.date
}

func (e *Image) Info() info.Info {
	return e.info
}

func (e *Image) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.Title(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Image) Section() string {
	return e.Parent().Section()
}

func (e *Image) Perma(lang string) string {
	return fmt.Sprintf("%v--not-implemented--%v", e.parent.Perma(lang), e.Slug(lang))
}
