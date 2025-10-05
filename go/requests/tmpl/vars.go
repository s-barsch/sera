package tmpl

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Paths struct {
	Root string
	Data string
}

type Config struct {
	FrontSettings *FrontSettings
	Strings       map[string]string
	Inlines       map[string]string
	Paths         Paths
}

type Article struct {
	TitleDe string `yaml:"title"`
	TitleEn string `yaml:"title-en"`
	Hash    string `yaml:"hash"`
}

func (a *Article) Title(lang string) string {
	if lang == "de" {
		return a.TitleDe
	}
	return a.TitleEn
}

type FrontSettings struct {
	Graph    int    `yaml:"graph-num"`
	Index    int    `yaml:"indecs-num"`
	Log      int    `yaml:"log-num"`
	Featured string `yaml:"featured"`
	Articles []*Article
}

func (v Config) Lang(key, lang string) string {
	return v.Strings[fmt.Sprintf("%v-%v", strings.ToLower(key), lang)]
}

func LoadVars(root string) (*Config, error) {
	s, err := ReadVarFiles(root)
	if err != nil {
		return nil, err
	}

	js, css, err := ReadVideoMainFiles(root)
	if err != nil {
		return nil, err
	}
	s["video-main.css"] = css
	s["video-main.js"] = js

	modtime, err := JsModtime(root)
	if err != nil {
		return nil, err
	}

	s["jsmodtime"] = modtime

	fr, err := ReadFrontSettings(root)
	if err != nil {
		return nil, err
	}

	inlines, err := ReadInlineStatics(root)
	if err != nil {
		return nil, err
	}

	return &Config{
		Strings:       s,
		Inlines:       inlines,
		FrontSettings: fr,
	}, nil
}

func ReadFrontSettings(root string) (*FrontSettings, error) {
	path := root + "/data/front/front.txt"
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fr := &FrontSettings{}
	err = yaml.Unmarshal(b, &fr)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func ReadVarFiles(root string) (map[string]string, error) {
	vars := map[string]string{}
	for _, name := range []string{"descriptions", "headings"} {

		path := fmt.Sprintf("/html/vars/%v.txt", name)
		b, err := os.ReadFile(root + path)
		if err != nil {
			return nil, err
		}

		m := map[string]string{}
		err = yaml.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}

		for k, v := range m {
			if vars[k] != "" {
				return nil, fmt.Errorf("duplicate entry in Vars: %v", k)
			}
			vars[k] = v
		}
	}
	return vars, nil
}

func ReadInlineStatics(root string) (map[string]string, error) {
	inlines := map[string]string{}

	sources := map[string]string{
		"css": "/css/dist/main.css",
		/*
			"indecsmap-de": "/static/svg/indecsmap-de.svg",
			"indecsmap-en": "/static/svg/indecsmap-en.svg",
		*/
		"logo-mobile": "/static/svg/logo/sferal-c-m.svg",
		"logo-desk":   "/static/svg/logo/sferal-c.svg",
		"pine":        "/static/img/icons/pine-c.svg",
		"email":       "/static/svg/email.svg",
	}
	for name, path := range sources {
		content, err := os.ReadFile(root + path)
		if err != nil {
			return nil, err
		}
		inlines[name] = string(content)
	}

	for i := 2007; i <= 2021; i++ {
		year := strconv.Itoa(i)
		path := fmt.Sprintf(root+"/static/svg/years/%v.svg", year)
		content, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		inlines[year] = string(content)
	}
	return inlines, nil
}
