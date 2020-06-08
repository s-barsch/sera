package tmpl

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

type Vars struct {
	FrontSettings *FrontSettings
	Strings map[string]string
}

type FrontSettings struct {
	Graph int `yaml:"graph-num"`
	Index int `yaml:"index-num"`
}

func (v Vars) Lang(key, lang string) string {
	return v.Strings[fmt.Sprintf("%v-%v", strings.ToLower(key), lang)]
}

func LoadVars(root string) (*Vars, error) {
	s, err := ReadVarFiles(root)
	if err != nil {
		return nil, err
	}

	modtime, err := JsModtime(root)
	if err != nil {
		return nil, err
	}

	maps, err := ReadIndexmap(root)
	if err != nil {
		return nil, err
	}

	logo, err := ReadLogo(root)
	if err != nil {
		return nil, err
	}

	css, err := ReadCSS(root)
	if err != nil {
		return nil, err
	}

	s["jsmodtime"] = modtime
	s["indexmap-de"] = maps["de"]
	s["indexmap-en"] = maps["en"]
	s["logo"] = logo
	s["css"] = css

	fr, err := ReadFrontSettings(root)
	if err != nil {
		return nil, err
	}

	return &Vars{
		Strings: s,
		FrontSettings: fr,
	}, nil
}

func ReadFrontSettings(root string) (*FrontSettings, error) {
	path := fmt.Sprintf(root + "/data/front/front.txt")
	b, err := ioutil.ReadFile(path)
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
	for _, name := range []string{"descriptions", "links", "headings"} {

		path := fmt.Sprintf("/html/vars/%v.txt", name)
		b, err := ioutil.ReadFile(root + path)
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
				return nil, fmt.Errorf("Duplicate entry in Vars: %v", k)
			}
			vars[k] = v
		}
	}
	return vars, nil
}
