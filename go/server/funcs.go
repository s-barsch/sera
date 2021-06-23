package server

import (
	"sacer/go/entry"
	"sacer/go/entry/tools"
	"sacer/go/entry/types/set"
	"sacer/go/entry/types/audio"
	"sacer/go/entry/types/video"
	"sacer/go/server/tmpl"
	"strings"
	"text/template"
	"time"
)

func (s *Server) Funcs() template.FuncMap {
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
			return s.Vars.FrontSettings.Articles
		},
		"var": func(name, lang string) string {
			return s.Vars.Lang(name, lang)
		},
		"varRaw": func(name string) string {
			return s.Vars.Strings[name]
		},
		"inlineFile": func(name string) string {
			return s.Vars.Inlines[name]
		},
		"inlineFileLang": func(name, lang string) string {
			return s.Vars.Inlines[name+"-"+lang]
		},
		"isLocal": func() bool {
			return s.Flags.Local
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
			return hasCaptions(s, lang)
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
		"hasCaptions": func(e entry.Entry, lang string) bool {
			s, ok := e.(*set.Set)
			if !ok {
				return false
			}
			return hasCaptions(s, lang)
		},
		"hasTranscript": func(e entry.Entry, lang string) bool {
			s, ok := e.(*set.Set)
			if !ok {
				return false
			}
			return hasTranscript(s, lang)
		},
		"nL":   tmpl.NewNotesLang,
		"eL":   tmpl.NewEntryLang,
		"eLy":  tmpl.NewEntryLangLazy,
		"esL":  tmpl.NewEntriesLang,
		"esLy": tmpl.NewEntriesLangLazy,
		"snav": tmpl.NewSubnav,

		"shaveParagraph": tools.ShaveParagraph,
		"minifySvg":      minifySVG,
	}
}

func hasCaptions(s *set.Set, lang string) bool {
	for _, child := range s.Entries() {
		m, ok := child.(entry.Media)
		if ok {
			return m.HasCaptions(lang)
		}
	}
	return false
}

func hasTranscript(s *set.Set, lang string) bool {
	for _, child := range s.Entries() {
		v, ok := child.(*video.Video)
		if ok {
			return v.Info().Field("transcript", lang) != ""
		}
		a, ok := child.(*audio.Audio)
		if ok {
			return a.Info().Field("transcript", lang) != ""
		}
	}
	return false
}
