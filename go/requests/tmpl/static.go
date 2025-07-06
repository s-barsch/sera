package tmpl

import (
	"fmt"
	"os"
	"path/filepath"
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

/*
func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}

func minifyCSS(str string) (string, error) {
	m := minify.New()
	m.Add("text/css", &css.Minifier{})
	return m.String("text/css", str)
}
*/

func ReadVideoMainFiles(root string) (js string, css string, err error) {
	path := "/static/js/video"
	l, nerr := os.ReadDir(root + path)
	if nerr != nil {
		err = nerr
		return
	}
	for _, fi := range l {
		if len(fi.Name()) > 4 && fi.Name()[:4] == "main" {
			if ext := filepath.Ext(fi.Name()); ext == ".js" {
				js = filepath.Join(path, fi.Name())
				continue
			} else if ext == ".css" {
				css = filepath.Join(path, fi.Name())
			}
		}
	}
	if css == "" || js == "" {
		err = fmt.Errorf("could not find video main js AND css.\ncss: %v\njs: %v", css, js)
		return
	}
	return
}
