package index 

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"st/pkg/el"
	"st/pkg/head"
	"st/pkg/server"
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
		Tree *el.Hold
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

