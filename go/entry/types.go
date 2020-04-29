package entry

import (
	//"reflect"
	"time"
)

//var Debug bool


/*
func Type(e interface{}) string {
	return reflect.TypeOf(e).String()
}
*/

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

type Video struct {
	File *File

	Date time.Time
	Info Info

	Subtitles []string
}
