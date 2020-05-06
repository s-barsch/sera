package server

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"stferal/go/entry"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/tree"
	"strings"
	"text/template"
	"time"
)

type objArg struct {
	Obj  interface{}
	Lang string
}

type entryLangObject struct {
	Entry entry.Entry
	Lazy  bool
	Lang  string
}

func (e *entryLangObject) E() entry.Entry {
	return e.Entry
}

func (e *entryLangObject) L() string {
	return e.Lang
}


/*
type elArg struct {
	El   interface{}
	Lazy bool
	Lang string
}

type elsArg struct {
	Els  entry.Els
	Lazy bool
	Lang string
}

type holdArg struct {
	Hold *entry.Hold
	Lazy bool
	Lang string
}

*/
type subnavObject struct {
	Tree    *tree.Tree
	Active  int64
	Lang    string
}

func (s *subnavObject) T() *tree.Tree {
	return s.Tree
}

func (s *subnavObject) L() string {
	return s.Lang
}

func (s *subnavObject) NavTrees() tree.Trees {
	t := s.Tree
	if t.Section() == "graph" {
		if t.Level() == 0 {
			return t.Trees.Reverse()
		}
		// if only one month
		if len(t.Trees) < 2 {
			return nil
		}
	}
	return t.Trees
}

func (s *subnavObject) IsYear() bool {
	return s.Tree.Level() == 0 && s.Tree.Section() == "graph"
}

var years = map[string]string{
	"de": "Jahre",
	"en": "Years",
}

func (s *subnavObject) YearLabel(lang string) string {
	return years[lang]
}
	

func (s *Server) TemplateFuncs() template.FuncMap {
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
		"monthLang": func(t time.Time, lang string) string {
			return helper.MonthLang(t, lang)
		},
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
		"eL": func(e entry.Entry, lang string) *entryLangObject {
			return &entryLangObject{
				Entry: e,
				Lang:  lang,
			}
		},
		"eLy": func(e entry.Entry, lazy bool, lang string) *entryLangObject {
			return &entryLangObject{
				Entry: e,
				Lazy: lazy,
				Lang:  lang,
			}
		},
		"snav": func(tree *tree.Tree, active int64, lang string) *subnavObject {
			return &subnavObject{
				Tree:   tree,
				Active: active,
				Lang:   lang,
			}
		},
		"minifySvg": minifySVG,
	}
}
