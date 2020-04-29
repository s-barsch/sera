package entry

import (
	//"reflect"
	"time"
)

//var Debug bool

const Timestamp = "060102_150405"

/*
func Type(e interface{}) string {
	return reflect.TypeOf(e).String()
}
*/

type Args []*Arg

type Arg struct {
	El   interface{}
	Lang string
}

func (args Args) Offset(start, end int) Args {
	l := len(args)
	if l < start {
		return Args{}
	}
	if end > l || end <= 0 {
		return args[start:]
	}
	return args[start:end]
}

type Els []interface{}

type Set struct {
	File *File

	Date  time.Time
	Info  Info
	Cover *Image

	Els Els
}

type Text struct {
	File *File

	Date time.Time
	Info Info

	Text  map[string]string
	Blank map[string]string
}

type Html struct {
	File *File

	Date time.Time
	Info Info

	Html map[string]string
}

type Image struct {
	File *File

	Date time.Time
	Dims *dims
	Info Info
}

type Audio struct {
	File *File

	Date time.Time
	Info Info

	Subtitles []string
}

type Video struct {
	File *File

	Date time.Time
	Info Info

	Subtitles []string
}
