package server

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

func (s *Server) loadRender() error {
	err := s.loadVars()
	if err != nil {
		return err
	}

	err = s.loadTemplates()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) loadTemplates() error {
	t := template.New("").Funcs(s.TemplateFuncs())

	t, err := t.ParseGlob(s.Paths.Root + "/html/*/*.html")
	if err != nil {
		return err
	}
	t, err = t.ParseGlob(s.Paths.Root + "/html/*/*/*.html")
	if err != nil {
		return err
	}
	s.Templates = t

	return nil
}

type vars map[string]string

func (v vars) Lang(key, lang string) string {
	return v[fmt.Sprintf("%v-%v", strings.ToLower(key), lang)]
}

func (s *Server) loadVars() error {
	v, err := s.readVarFiles()
	if err != nil {
		return err
	}

	modtime, err := s.readJsModtime()
	if err != nil {
		return err
	}

	maps, err := s.readIndexmap()
	if err != nil {
		return err
	}

	logo, err := s.readLogo()
	if err != nil {
		return err
	}

	css, err := s.readCSS()
	if err != nil {
		return err
	}
	v["jsmodtime"] = modtime
	v["indexmap-de"] = maps["de"]
	v["indexmap-en"] = maps["en"]
	v["logo"] = logo
	v["css"] = css

	s.Vars = v

	return nil
}

func (s *Server) readVarFiles() (map[string]string, error) {
	vars := map[string]string{}
	for _, name := range []string{"descriptions", "links", "headings"} {

		path := fmt.Sprintf("/html/vars/%v.txt", name)
		b, err := ioutil.ReadFile(s.Paths.Root + path)
		if err != nil {
			return nil, err
		}

		m := map[string]string{}
		err = yaml.Unmarshal([]byte(b), &m)
		if err != nil {
			return nil, err
		}

		for k, v := range m {
			if vars[k] != "" {
				return nil, fmt.Errorf("Duplicate entry in Vars: %v", k)
			}
			vars[k] = v
		}
	}
	return vars, nil
}

func makeTimestamp(t time.Time) string {
	s := fmt.Sprintf("%x", t.Unix())
	const length = 3
	if len(s) > length {
		return s[len(s)-length:]
	}
	return s
}

func (s *Server) readJsModtime() (string, error) {
	fi, err := os.Stat(s.Paths.Data + "/static/js/bundle.js")
	if err != nil {
		return "", err
	}
	return makeTimestamp(fi.ModTime()), nil
}

func (s *Server) readLogo() (string, error) {
	return readFile(s.Paths.Data + "/static/svg/stferal-logo-compressed.svg")
}

func (s *Server) readIndexmap() (map[string]string, error) {
	m := map[string]string{}
	for _, lang := range []string{"de", "en"} {
		path := fmt.Sprintf(s.Paths.Data+"/static/svg/indexmap-%v.svg", lang)
		str, err := readFile(path)
		if err != nil {
			return nil, err
		}
		m[lang] = str
	}
	return m, nil
}

func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}

func (s *Server) readCSS() (string, error) {
	b, err := ioutil.ReadFile(s.Paths.Root + "/css/dist/main.css")
	if err != nil {
		return "", err
	}

	return minifyCSS(string(b))
}
