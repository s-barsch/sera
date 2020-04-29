// +build ignore

package entry

func NewEl(path string, hold *Hold) (interface{}, error) {
	switch FileType(path) {
	case "image":
		return NewImage(path, hold)
	case "audio":
		return NewAudio(path, hold)
	case "video":
		return NewVideo(path, hold)
	case "text":
		return NewText(path, hold)
	case "dir":
		return NewSet(path, hold)
	case "html":
		return NewHtml(path, hold)
	default:
		return NewFile(path, hold)
	}
}
