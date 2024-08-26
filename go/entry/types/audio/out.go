package audio

import (
	"fmt"

	"g.sacerb.com/sacer/go/entry/tools"
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

func (a *Audio) CaptionLocation(lang string) string {
	return tools.VTTPath(a.file.Dir(), a.file.NameNoExt(), lang)
}

func (a *Audio) FilePath(lang string) string {
	switch a.parent.Type() {
	case "set", "tree":
		return fmt.Sprintf("%v/cache/%v", a.parent.Perma(lang), a.file.Name())
	}
	return fmt.Sprintf("%v/cache/%v", a.Perma(lang), a.file.Name())
}

func (a *Audio) HasCaptions(lang string) bool {
	for _, captionsLang := range a.Captions {
		if lang == captionsLang {
			return true
		}
	}
	return false
}

func (a *Audio) Captioned() bool {
	return len(a.Captions) == 2
}

func (a *Audio) Transcripted() bool {
	for _, str := range a.Transcript.Langs {
		if str == "" {
			return false
		}
	}
	return true
}
