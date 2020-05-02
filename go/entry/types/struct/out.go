package stru

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

func (s *Struct) Id() string {
	return s.date.Format(helper.Timestamp)
}

func (s *Struct) Date() time.Time {
	return s.date
}

func (s *Struct) Info() info.Info {
	return s.info
}

