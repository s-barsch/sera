package entry

import (
	"time"
	"stferal/go/entry/parts/info"
)

type Entry interface{
	Id()    string

	Info()  info.Info
	Date()  time.Time

	Title(string) string
	//Perma(string) string
}

type Entries []Entry
