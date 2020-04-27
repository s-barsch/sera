package server

import (
	"bytes"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	//"github.com/yosssi/gohtml"
	//"net/http"
	"io"
)

func (s *Server) ExecuteTemplate(w io.Writer, tmpl string, d interface{}) error {
	mw := MinifyFilter("text/html", w)
	defer mw.Close()
	w = mw
	return s.Templates.ExecuteTemplate(w, tmpl, d)
}

// Used within HTML templates.
func (s *Server) RenderTemplate(tname string, d interface{}) (string, error) {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals: true,
		KeepWhitespace:      true,
		//KeepEndTags:         true,
	})
	var b bytes.Buffer
	err := s.Templates.ExecuteTemplate(&b, tname, d)
	if err != nil {
		return "", err
	}
	return m.String("text/html", b.String())
}
