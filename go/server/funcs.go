package server

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"stferal/go/server/tmpl"
	"stferal/go/entry/helper"
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
		"removeß": func(str string) string {
			return strings.Replace(str, "ß", "ss", -1)
		},
		"hyphen": func(str, lang string) (string, error) {
			return "", fmt.Errorf("Template hyphen function currently not implemented.")
		},
		"normalize": func(str string) string {
			return helper.Normalize(str)
		},
		"executeTemplate": func(name string, data interface{}) string {
			buf := &bytes.Buffer{}
			if err := s.Templates.ExecuteTemplate(buf, name, data); err != nil {
				log.Println(err)
				return "template err"
			}
			return buf.String()
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
		/*
		"lastEl": func(els entry.Els) interface{} {
			if len(els) <= 0 {
				return nil
			}
			return els[len(els)-1]
		},
		"makeGraphMoreLink": func(year, lang string) (string, error) {
			str := s.Vars.Lang("graph-main-more", lang)
			href := fmt.Sprintf("/graph/%v", year)
			return fmt.Sprintf(str, href), nil
		},
		*/
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
		"filepathDir": filepath.Dir,
		"title":       strings.Title,
		"upper":       strings.ToUpper,
		"tolower":     strings.ToLower,
		"esc":         template.HTMLEscapeString,
		"render":      s.RenderTemplate,
		"monthLang": helper.MonthLang,
		"nodeName": func(acr string) string {
			return "node_" + acr
		},
		"min1": func(x int) int {
			return x - 1
		},
		"plus1": func(x int) int {
			return x + 1
		},
		"iso8601": func(date time.Time) string {
			return date.Format(time.RFC3339)
		},
		"eL": tmpl.NewEntryLang,
		"eLy": tmpl.NewEntryLangLazy,
		"esL": tmpl.NewEntriesLang,
		"esLy": tmpl.NewEntriesLangLazy,
		"snav": tmpl.NewSubnav,
		"minifySvg": minifySVG,
	}
}
