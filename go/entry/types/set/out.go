package set

import (
	"stferal/go/entry/helper"
	"stferal/go/entry/parts/info"
	"time"
)

func (s *Set) Id() string {
	return s.date.Format(helper.Timestamp)
}

func (s *Set) Date() time.Time {
	return s.date
}

func (s *Set) Info() info.Info {
	return s.info
}

