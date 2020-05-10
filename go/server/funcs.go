package server

import (
	"stferal/go/entry/helper"
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
		"var": func(name, lang string) string {
			return s.Vars.Lang(name, lang)
		},
		"varRaw": func(name string) string {
			return s.Vars[name]
		},
		"isLocal": func() bool {
			return s.Flags.Local
		},
		"langName": func(lang string) string {
			return helper.LangNames[lang]
		},
		"monthLang":   helper.MonthLang,
		"nodeName": func(id int64) string {
			return "node_" + helper.ToTimestamp(id)
		},
		"plus1": func(x int) int {
			return x + 1
		},
		"iso8601": func(date time.Time) string {
			return date.Format(time.RFC3339)
		},
		"eL":        tmpl.NewEntryLang,
		"eLy":       tmpl.NewEntryLangLazy,
		"esL":       tmpl.NewEntriesLang,
		"esLy":      tmpl.NewEntriesLangLazy,
		"snav":      tmpl.NewSubnav,
		"minifySvg": minifySVG,
	}
}
