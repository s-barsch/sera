package image

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

func (i *Image) Id() string {
	return i.date.Format(helper.Timestamp)
}

func (i *Image) Date() time.Time {
	return i.date
}

func (i *Image) Info() info.Info {
	return i.info
}
