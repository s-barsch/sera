package transcript

import (
	"fmt"
	"os"
	"path/filepath"

	"g.rg-s.com/sera/go/entry/info"
	"g.rg-s.com/sera/go/entry/tools"
	"g.rg-s.com/sera/go/entry/tools/script"
)

func GetTranscripts(p string) (*script.Script, error) {
	langMap, err := findTranscript(p)
	if err != nil {
		return nil, err
	}

	if langMap == nil {
		return script.EmptyScript(), nil
	}

	script := script.RenderScript(langMap)
	script.NumberFootnotes(1)

	return script, nil
}

func findTranscript(p string) (script.LangMap, error) {
	dir := filepath.Dir(p)
	base := tools.StripExt(filepath.Base(p))
	langs := script.LangMap{}
	found := false
	for lang := range tools.Langs {
		name := fmt.Sprintf("%v.%v.txt", base, lang)
		path := filepath.Join(dir, "transcript", name)
		_, err := os.Stat(path)
		if err != nil {
			continue
		}
		f, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		langs[lang] = string(f)
		found = true
	}
	if !found {
		return findTranscriptInfo(p)
	}
	return langs, nil
}

func findTranscriptInfo(path string) (script.LangMap, error) {
	if info.HasFileInfo(path) {
		inf, err := info.ReadFileInfo(path)
		if err != nil {
			return nil, err
		}
		return extractTranscript(inf), nil
	}
	return nil, nil
}

func extractTranscript(i info.Info) script.LangMap {
	langs := script.LangMap{}
	for l := range tools.Langs {
		key := "transcript"
		if l != "de" {
			key += "-" + l
		}
		langs[l] = i[key]
	}
	return langs
}
