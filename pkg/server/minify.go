package server

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"io"
)

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

/*
func minifyHTML(w http.ResponseWriter) MinifyResponseWriter {
	return MinifyFilter("text/html", w)
}

func minifyHTML(str string) (string, error) {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals: true,
		KeepWhitespace: true,
	})
	return m.String("text/html", str)
}
*/

func minifySVG(str string) (string, error) {
	m := minify.New()
	m.Add("text/svg", &svg.Minifier{
		Decimals: -1,
	})
	return m.String("text/svg", str)
}

func minifyCSS(str string) (string, error) {
	m := minify.New()
	m.Add("text/css", &css.Minifier{
		Decimals: -1,
	})
	return m.String("text/css", str)
}
