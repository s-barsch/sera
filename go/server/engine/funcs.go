package engine

import (
	"fmt"
	"math"
	"strings"
	"text/template"
	"time"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/entry/tools"
	"g.rg-s.com/sacer/go/entry/types/set"
	"g.rg-s.com/sacer/go/entry/types/video"
	"g.rg-s.com/sacer/go/requests/tmpl"
)

func (e *Engine) Funcs() template.FuncMap {
	return template.FuncMap{
		"aboutTitle": func(lang string) string {
			switch lang {
			case "de":
				return "Über"
			}
			return "About"
		},
		"normalize": tools.Normalize,
		"removeß": func(str string) string {
			return strings.Replace(str, "ß", "ss", -1)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"rel": func(path string) string {
			x := len(s.Paths.Data)
			if len(path) > x {
				return path[x:]
			}
			return path
		},
		"abbr": tools.Abbr,
		"frontArticles": func() []*tmpl.Article {
			return s.Engine.Vars.FrontSettings.Articles
		},
		"var": func(name, lang string) string {
			return s.Engine.Vars.Lang(name, lang)
		},
		"varRaw": func(name string) string {
			return s.Engine.Vars.Strings[name]
		},
		"inlineFile": func(name string) string {
			return s.Engine.Vars.Inlines[name]
		},
		"inlineFileLang": func(name, lang string) string {
			return s.Engine.Vars.Inlines[name+"-"+lang]
		},
		"isLocal": func() bool {
			return s.Flags.Local
		},
		"displayInfo": func() bool {
			return s.Flags.Info
		},
		"langName": func(lang string) string {
			return tools.Langs[lang]
		},
		"monthLang": tools.MonthLang,
		"nodeName": func(id int64) string {
			return "node_" + tools.ToTimestamp(id)
		},
		"plus1": func(x int) int {
			return x + 1
		},
		"iso8601": func(date time.Time) string {
			return date.Format(time.RFC3339)
		},
		"isTranslated": func(e entry.Entry, lang string) bool {
			if lang == "de" {
				return true
			}
			if x := e.Info()["translated"]; x != "" {
				return x == "true"
			}
			s, ok := e.(*set.Set)
			if !ok {
				return false
			}
			return isCaptioned(s)
		},
		"setVideo": func(s *set.Set) *video.Video {
			for _, child := range s.Entries() {
				v, ok := child.(*video.Video)
				if ok {
					return v
				}
			}
			return nil
		},
		"duration": func(s *set.Set) (string, error) {
			for _, child := range s.Entries() {
				v, ok := child.(*video.Video)
				if ok {
					mins := (v.Duration / 60) - 1
					secs := math.Mod(v.Duration, 60)
					return fmt.Sprintf("%.0f:%.0f", mins, secs), nil
				}
			}
			fmt.Println(s)
			return "", fmt.Errorf("err 'duration': has no entries")
		},
		"isCaptioned":    isCaptioned,
		"isTranscripted": isTranscripted,
		"nL":             tmpl.NewNotesLang,
		"eL":             tmpl.NewEntryLang,
		"eLy":            tmpl.NewEntryLangLazy,
		"esL":            tmpl.NewEntriesLang,
		"esLy":           tmpl.NewEntriesLangLazy,
		"snav":           tmpl.NewSubnav,

		"shaveParagraph": tools.ShaveParagraph,
		"minifySvg":      minifySVG,
	}
}

func isCaptioned(s *set.Set) bool {
	for _, child := range s.Entries() {
		m, ok := child.(entry.Media)
		if ok {
			return m.Captioned()
		}
	}
	return false
}

func isTranscripted(s *set.Set) bool {
	for _, child := range s.Entries() {
		m, ok := child.(entry.Media)
		if ok {
			return m.Transcripted()
		}
	}
	return false
}
