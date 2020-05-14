package video

import (
	"fmt"
)

func (v *Video) Location(arg string) string {
	return v.file.Path
}

func (v *Video) FilePath(lang string) string {
	if v.parent.Type() == "set" {
		return fmt.Sprintf("%v/cache/%v", v.parent.Perma(lang), v.file.Name())
	}
	return fmt.Sprintf("%v/cache/%v", v.Perma(lang), v.file.Name())
}


func (v *Video) SubtitlePath(lang string) string {
	return fmt.Sprintf(
		"%v/files/vtt/%v-%v.vtt",
		v.parent.Perma(lang),
		v.file.NameNoExt(),
		lang,
	)
}


