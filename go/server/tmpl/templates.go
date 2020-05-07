package tmpl

import (
	"text/template"
)

func LoadTemplates(root string, funcs template.FuncMap) (*template.Template, error) {
	t := template.New("").Funcs(funcs)

	t, err := t.ParseGlob(root + "/html/*/*.html")
	if err != nil {
		return nil, err
	}
	return t.ParseGlob(root + "/html/*/*/*.html")
}


