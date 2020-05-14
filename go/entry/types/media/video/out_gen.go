// Code generated by go generate; DO NOT EDIT.

package video

import (
	"fmt"

	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Video) Type() string {
	return "video"
}

func (e *Video) Parent() entry.Entry {
	return e.parent
}

func (e *Video) File() *file.File {
	return e.file
}

func (e *Video) Id() int64 {
	return e.date.Unix()
}

func (e *Video) Timestamp() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Video) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Video) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Video) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Video) Date() time.Time {
	return e.date
}

func (e *Video) Info() info.Info {
	return e.info
}

func (e *Video) Slug(lang string) string {
	if slug := e.info.Slug(lang); slug != "" {
		return slug
	}
	return helper.Normalize(e.info.Title(lang))
}

func (e *Video) IsBlob() bool {
	return entry.IsBlob(e)
}

func (e *Video) SetParent(parent entry.Entry) {
	e.parent = parent
}

func (e *Video) SetInfo(inf info.Info) {
	e.info = inf
}

func (e *Video) Path(lang string) string {
	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Slug(lang))
}

// This recursive function call will be caught by a Tree type. For now, all
// further up parent entries are exclusively of type Tree.
func (e *Video) Section() string {
	return e.Parent().Section()
}

func (e *Video) Perma(lang string) string {
	if e.Section() == "index" {
		return fmt.Sprintf("%v#%v", e.parent.Perma(lang), helper.Normalize(e.Title(lang)))
	}
	slug := e.Slug(lang)
	if slug != "" {
		return fmt.Sprintf("%v/%v-%v", e.parent.Path(lang), slug, e.Hash())
	}

	return fmt.Sprintf("%v/%v", e.parent.Path(lang), e.Hash())
}