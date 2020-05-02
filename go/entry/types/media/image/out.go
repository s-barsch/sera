package image

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

func (e *Image) Id() string {
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
