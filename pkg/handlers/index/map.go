package index

import (
	"bytes"
	"bufio"
	"fmt"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/svg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"stferal/pkg/entry"
	"stferal/pkg/head"
	"stferal/pkg/server"
)

func MapDot(s *server.Server, w http.ResponseWriter, r *http.Request) {
	err := printMapDot(s, w, head.Lang(r.Host))
	if err != nil {
		log.Println(err)
	}
}

func printMapDot(s *server.Server, w io.Writer, lang string) error {
	return s.Templates.ExecuteTemplate(w, "map", struct {
		Lang string
		Tree *entry.Hold
	}{
		Lang: lang,
		Tree: s.Trees["index"],
	})
}

func MapSVG(s *server.Server, w http.ResponseWriter, r *http.Request) {
	b, err := renderMap(s, head.Lang(r.Host), "svg")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "%s", b)
}

func renderMap(s *server.Server, lang string, filetype string) ([]byte, error) {
	buf := bytes.Buffer{}
	err := printMapDot(s, &buf, lang)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("neato", fmt.Sprintf("-T%v", filetype))
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, buf.String())
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func SaveMaps(s *server.Server) error {
	for _, lang := []string{"en", "de"} {
		err := saveMap(s, lang)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveMap(s *server.Server, lang string) error {
	b, err := renderMap(s, lang, "svg")
	if err != nil {
		return err
	}

	minified, err := optimizeSVG(b)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(s.Paths.Data + "/static/svg/indexmap-%v.svg", lang)

	err = ioutil.WriteFile(path, minified, 0644)
	if err != nil {
		return err
	}
	return nil
}

func optimizeSVG(body []byte) ([]byte, error) {
	var err error
	body, err = stripTitles(body)
	if err != nil {
		return nil, err
	}

	return minifySVG(body)
}

func stripTitles(input []byte) ([]byte, error) {
	buf := bytes.Buffer{}
	s := bufio.NewScanner(bytes.NewReader(input))
	for s.Scan() {
		text := s.Text()
		if len(text) >= 7 && text[:4] == "<svg" {
			text = `<svg class="indexmap"` + text[4:]
		}
		if len(text) >= 7 && text[:7] == "<title>" {
			continue
		}
		buf.WriteString(text)
	}
	return buf.Bytes(), nil
}

func minifySVG(input []byte) ([]byte, error) {
	m := minify.New()
	m.AddFunc("image/svg+xml", svg.Minify)
	return m.Bytes("image/svg+xml", input)
}

