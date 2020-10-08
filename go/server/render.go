package server

import (
	"bytes"
	"io"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
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
		//KeepEndTags:       true,
	})
	var b bytes.Buffer
	err := s.Templates.ExecuteTemplate(&b, tname, d)
	if err != nil {
		return "", err
	}
	return m.String("text/html", b.String())
}

// minify

type MinifyResponseWriter struct {
	io.Writer
	io.WriteCloser
}

func (m MinifyResponseWriter) Write(b []byte) (int, error) {
	return m.WriteCloser.Write(b)
}

// MinifyResponseWriter must be closed explicitly by calling site.
func MinifyFilter(mediatype string, res io.Writer) MinifyResponseWriter {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals: true,
		KeepWhitespace:      true,
		//KeepDocumentTags: true,
		//KeepEndTags: true,
	})

	m.Add("application/ld+json", &json.Minifier{})
	/*
		if !s.Flags.Debug {
			m.Add("application/ld+json", &json.Minifier{})
		}
	*/

	mw := m.Writer(mediatype, res)
	return MinifyResponseWriter{res, mw}
}

func minifySVG(str string) (string, error) {
	m := minify.New()
	m.Add("text/svg", &svg.Minifier{
		Decimals: -1,
	})
	return m.String("text/svg", str)
}
