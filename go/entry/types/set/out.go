package set

import (
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/file"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Set) Parent() entry.Entry {
	return e.parent
}

func (e *Set) File() *file.File {
	return e.file
}

func (e *Set) Id() string {
	return e.date.Format(helper.Timestamp)
}

func (e *Set) Hash() string {
	return helper.ToB16(e.date)
}

func (e *Set) HashShort() string {
	return helper.ShortenHash(e.Hash())
}

func (e *Set) Title(lang string) string {
	if title := e.info.Title(lang); title != "" {
		return title
	}
	return e.HashShort()
}

func (e *Set) Date() time.Time {
	return e.date
}

func (e *Set) Info() info.Info {
	return e.info
}
