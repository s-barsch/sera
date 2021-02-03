package tmpl

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

func JsModtime(root string) (string, error) {
	fi, err := os.Stat(root + "/static/js/bundle.js")
	if err != nil {
		return "", err
	}
	return makeTimestamp(fi.ModTime()), nil
}

func makeTimestamp(t time.Time) string {
	s := fmt.Sprintf("%x", t.Unix())
	const length = 3
	if len(s) > length {
		return s[len(s)-length:]
	}
	return s
}

func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}

func minifyCSS(str string) (string, error) {
	m := minify.New()
	m.Add("text/css", &css.Minifier{})
	return m.String("text/css", str)
}
