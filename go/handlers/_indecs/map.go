package indecs

import (
	//"bufio"
	"bytes"
	"fmt"

	//"github.com/tdewolff/minify"
	//"github.com/tdewolff/minify/svg"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"g.sacerb.com/sacer/go/entry/types/tree"
	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/meta"
)

func MapDot(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	err := printMapDot(s, w, m.Lang, false)
	if err != nil {
		log.Println(err)
	}
}

func printMapDot(s *server.Server, w io.Writer, lang string, all bool) error {
	t := s.Trees["indecs"].Access(false)[lang].Public()
	return s.Templates.ExecuteTemplate(w, "map", struct {
		Lang string
		Tree *tree.Tree
		All  bool
	}{
		Lang: lang,
		Tree: t,
		All:  all,
	})
}

func MapAll(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	buf := bytes.Buffer{}
	err := printMapDot(s, &buf, m.Lang, true)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b, err := renderMap(buf.String(), "svg")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "%s", b)
}

func MapIndex(s *server.Server, w http.ResponseWriter, r *http.Request, m *meta.Meta) {
	buf := bytes.Buffer{}
	err := printMapDot(s, &buf, m.Lang, false)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	b, err := renderMap(buf.String(), "svg")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "%s", b)
}

func renderMap(markup, filetype string) ([]byte, error) {
	cmd := exec.Command("neato", fmt.Sprintf("-T%v", filetype))
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, markup)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

// still needed?

/*
func SaveMaps(s *server.Server) error {
	for _, lang := range []string{"en", "de"} {
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

	path := fmt.Sprintf(s.Paths.Data+"/static/svg/indecsmap-%v.svg", lang)

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
			text = `<svg class="indecsmap"` + text[4:]
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
*/
