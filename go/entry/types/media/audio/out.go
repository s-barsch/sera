package audio

import (
	"fmt"
)

func (a *Audio) Location(arg string) string {
	return a.file.Path
}



func (a *Audio) CaptionPath(lang string) string {
	return fmt.Sprintf(
		"%v/files/vtt/%v-%v.vtt",
		a.parent.Perma(lang),
		a.file.NameNoExt(),
		lang,
	)
}

func (a *Audio) FilePath(lang string) string {
	if a.parent.Type() == "set" {
		return fmt.Sprintf("%v/cache/%v", a.parent.Perma(lang), a.file.Name())
	}
	return fmt.Sprintf("%v/cache/%v", a.Perma(lang), a.file.Name())
}


