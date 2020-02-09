package server

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"stferal/pkg/entry"
	"strings"
	"text/template"
	"time"
)

type objArg struct {
	Obj  interface{}
	Lang string
}

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

type subNavArg struct {
	Hold    *entry.Hold
	Current string
	Lang    string
}

func (s *Server) TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"removeß": func(str string) string {
			return strings.Replace(str, "ß", "ss", -1)
		},
		"hyphen": func(str, lang string) (string, error) {
			return "", fmt.Errorf("Template hyphen function currently not implemented.")
		},
		"normalize": func(str string) string {
			return entry.Normalize(str)
		},
		"executeTemplate": func(name string, data interface{}) string {
			buf := &bytes.Buffer{}
			if err := s.Templates.ExecuteTemplate(buf, name, data); err != nil {
				log.Println(err)
				return "template err"
			}
			return buf.String()
		},
		"polyfillPath": func() string {
			if s.Flags.Mobile {
				return "http://192.168.1.56:3000/v2/polyfill.min.js?features=forEach,IntersectionObserver"
			}
			return "/js/polyfill.min.js"
		},
		"add": func(a, b int) int {
			return a + b
		},
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
			return entry.LangNames[lang]
		},
		"filepathDir": filepath.Dir,
		"title":       strings.Title,
		"upper":       strings.ToUpper,
		"tolower":     strings.ToLower,
		"esc":         template.HTMLEscapeString,
		"render":      s.RenderTemplate,
		"snavArg": func(hold *entry.Hold, current, lang string) *subNavArg {
			return &subNavArg{
				Hold:    hold,
				Current: current,
				Lang:    lang,
			}
		},
		"monthLang": func(t time.Time, lang string) string {
			return entry.MonthLang(t, lang)
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
		"holdArg": func(h *entry.Hold, lazy bool, lang string) *holdArg {
			return &holdArg{
				Hold: h,
				Lazy: lazy,
				Lang: lang,
			}
		},
		"elArg": func(e interface{}, lazy bool, lang string) *elArg {
			return &elArg{
				El:   e,
				Lazy: lazy,
				Lang: lang,
			}
		},
		"elsArg": func(els entry.Els, lazy bool, lang string) *elsArg {
			return &elsArg{
				Els:  els,
				Lazy: lazy,
				Lang: lang,
			}
		},
		"objArg": func(obj interface{}, lang string) *objArg {
			return &objArg{
				Obj:  obj,
				Lang: lang,
			}
		},
		"elType":    entry.Type,
		"minifySvg": minifySVG,
	}
}
