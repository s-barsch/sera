package media

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/entry/helper"
	"sacer/go/entry/types/media/audio"
	"sacer/go/entry/types/media/html"
	"sacer/go/entry/types/media/image"
	"sacer/go/entry/types/media/text"
	"sacer/go/entry/types/media/video"
)

func NewMediaEntry(path string, parent entry.Entry) (entry.Entry, error) {
	switch helper.FileType(path) {
	case "text":
		return text.NewText(path, parent)
	case "image":
		return image.NewImage(path, parent)
	case "audio":
		return audio.NewAudio(path, parent)
	case "video":
		return video.NewVideo(path, parent)
	case "html":
		return html.NewHtml(path, parent)
	}
	return nil, &helper.Err{
		Path: path,
		Func: "NewMediaEntry",
		Err:  fmt.Errorf("invalid entry type: %v", helper.FileType(path)),
	}
}
