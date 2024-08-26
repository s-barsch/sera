package media

import (
	"fmt"

	"g.sacerb.com/sacer/go/entry"
	"g.sacerb.com/sacer/go/entry/tools"
	"g.sacerb.com/sacer/go/entry/types/audio"
	"g.sacerb.com/sacer/go/entry/types/html"
	"g.sacerb.com/sacer/go/entry/types/image"
	"g.sacerb.com/sacer/go/entry/types/text"
	"g.sacerb.com/sacer/go/entry/types/video"
)

func NewMediaEntry(path string, parent entry.Entry) (entry.Entry, error) {
	switch tools.FileType(path) {
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
	return nil, &tools.Err{
		Path: path,
		Func: "NewMediaEntry",
		Err:  fmt.Errorf("invalid entry type: %v", tools.FileType(path)),
	}
}
