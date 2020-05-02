package text

import (
	"stferal/go/entry/helper"
	"time"
	"stferal/go/entry/parts/info"
)

func (t *Text) Id() string {
	return t.date.Format(helper.Timestamp)
}

func (t *Text) Date() time.Time {
	return t.date
}

func (t *Text) Info() info.Info {
	return t.info
}


