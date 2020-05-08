package tmpl

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Vars map[string]string

func (v Vars) Lang(key, lang string) string {
	return v[fmt.Sprintf("%v-%v", strings.ToLower(key), lang)]
}

func LoadVars(root string) (Vars, error) {
	v, err := ReadVarFiles(root)
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

	v["jsmodtime"] = modtime
	v["indexmap-de"] = maps["de"]
	v["indexmap-en"] = maps["en"]
	v["logo"] = logo
	v["css"] = css

	return v, nil
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
