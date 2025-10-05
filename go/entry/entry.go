package entry

import (
	"time"

	"g.rg-s.com/sacer/go/entry/file"
	"g.rg-s.com/sacer/go/entry/info"
)

type Entry interface {
	Parent() Entry
	File() *file.File

	Id() int64
	Type() string
	Section() string

	Hash() string
	Timestamp() string

	Info() info.Info
	Date() time.Time

	Title(string) string
	Path(string) string
	Perma(string) string

	MediaObject() bool
	ObjectType() string

	SetParent(Entry)
	SetInfo(info.Info)
}

type Collection interface {
	Entries() Entries
}

type Blob interface {
	Location(string, string) (string, error)
}

type Media interface {
	Transcripted() bool
	Captioned() bool
	HasCaptions(string) bool
}

func IsBlob(e Entry) bool {
	_, ok := e.(Blob)
	return ok
}
