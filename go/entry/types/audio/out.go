package audio

import (
	"fmt"
)

func (a *Audio) Location(arg string) (string, error) {
	return a.file.Path, nil
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
	switch a.parent.Type() {
	case "set", "tree":
		return fmt.Sprintf("%v/cache/%v", a.parent.Perma(lang), a.file.Name())
	}
	return fmt.Sprintf("%v/cache/%v", a.Perma(lang), a.file.Name())
}
