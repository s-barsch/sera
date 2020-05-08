package tmpl

import (
	"fmt"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"io/ioutil"
	"os"
	"time"
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

func ReadLogo(root string) (string, error) {
	return readFile(root + "/static/svg/stferal-logo-compressed.svg")
}

func ReadIndexmap(root string) (map[string]string, error) {
	m := map[string]string{}
	for _, lang := range []string{"de", "en"} {
		path := fmt.Sprintf(root+"/static/svg/indexmap-%v.svg", lang)
		str, err := readFile(path)
		if err != nil {
			return nil, err
		}
		m[lang] = str
	}
	return m, nil
}

func ReadCSS(root string) (string, error) {
	b, err := ioutil.ReadFile(root + "/css/dist/main.css")
	if err != nil {
		return "", err
	}

	return minifyCSS(string(b))
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
