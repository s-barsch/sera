package engine

import (
	"text/template"

	"g.rg-s.com/sacer/go/requests/tmpl"
)

type Engine struct {
	templates *template.Template
	Config    *tmpl.Config
}
