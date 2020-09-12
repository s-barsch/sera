package server

import (
	"stferal/go/entry/helper"
	"stferal/go/entry"
	"stferal/go/entry/types/media/video"
	"stferal/go/entry/types/set"
	"stferal/go/server/tmpl"
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
		"normalize": helper.Normalize,
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
		"abbr": helper.Abbr,
		"frontArticles": func() []*tmpl.Article {
			return s.Vars.FrontSettings.Articles
		},
		"var": func(name, lang string) string {
			return s.Vars.Lang(name, lang)
		},
		"varRaw": func(name string) string {
			return s.Vars.Strings[name]
		},
		"isLocal": func() bool {
			return s.Flags.Local
		},
		"langName": func(lang string) string {
			return helper.LangNames[lang]
		},
		"monthLang": helper.MonthLang,
		"nodeName": func(id int64) string {
			return "node_" + helper.ToTimestamp(id)
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
			return hasSubtitles(s, lang)
		},
		"eL":        tmpl.NewEntryLang,
		"eLy":       tmpl.NewEntryLangLazy,
		"esL":       tmpl.NewEntriesLang,
		"esLy":      tmpl.NewEntriesLangLazy,
		"snav":      tmpl.NewSubnav,
		"minifySvg": minifySVG,
	}
}

func hasSubtitles(s *set.Set, lang string) bool {
	for _, child := range s.Entries() {
		v, ok := child.(*video.Video)
		if ok {
			return v.HasSubtitles(lang)
		}
	}
	return false
}
