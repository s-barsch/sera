package meta

import (
	"fmt"
	"net/http"

	"g.rg-s.com/sacer/go/entry/tools"
)

type Options struct {
	Colors *Option
	Size   *Option
}

type Option struct {
	Name    string
	Values  []string
	Default string
	Active  string
}

type Value struct {
	Title    string
	Value    string
	Href     string
	IsActive bool
}

func (o *Option) List() []*Value {
	m := []*Value{}
	for _, v := range o.Values {
		m = append(m, &Value{
			Title:    valueTitle(o.Name, v),
			Value:    v,
			IsActive: v == o.Active,
			Href:     valueHref(o.Name, v),
		})
	}
	return m
}

func valueTitle(option, value string) string {
	if option == "size" {
		title := value
		if len(value) > 1 {
			title = value[:1]
		}
		return tools.Title(title)
	}
	return value
}

func GetOptions(r *http.Request) *Options {
	colors := &Option{
		Name: "colors",
		Values: []string{
			"light",
			"dark",
		},
		Default: "light",
	}
	colors.setActive(r)

	size := &Option{
		Name: "size",
		Values: []string{
			"small",
			"medium",
			"large",
		},
		Default: "medium",
	}
	size.setActive(r)

	return &Options{
		Colors: colors,
		Size:   size,
	}
}

func (o *Option) NextValue() string {
	c := 0
	for i, v := range o.Values {
		if v == o.Active {
			c = i + 1
			break
		}
	}
	if c >= len(o.Values) {
		return o.Values[0]
	}
	return o.Values[c]
}

func (o *Option) CurrentLink() string {
	return valueHref(o.Name, o.Active)
}

func (o *Option) NextLink() string {
	return valueHref(o.Name, o.NextValue())
}

func valueHref(option, value string) string {
	return fmt.Sprintf("/opt/%v/%v", option, value)
}

func (o *Option) setActive(r *http.Request) {
	c, _ := r.Cookie(o.Name)
	if c != nil && o.isValid(c.Value) {
		o.Active = c.Value
		return
	}
	o.Active = o.Default
}

func (o *Option) isValid(value string) bool {
	for _, v := range o.Values {
		if v == value {
			return true
		}
	}
	return false
}

/*
func (m *Meta) SwitchTypeTitle(lang string) string {
	switch lang {
	case "en":
		if h.Dark() {
			return "Switch to small type"
		} else {
			return "Switch to large type"
		}
	default:
		if h.Dark() {
			return "Wechsle zu großer Schrift"
		} else {
			return "Wechsle zu kleiner Schrift"
		}
	}
}

func (m *Meta) SwitchColorsTitle(lang string) string {
	switch lang {
	case "en":
		if h.Colors() == "dark" {
			return "Switch to light colors"
		} else {
			return "Switch to dark colors"
		}
	default:
		if h.Colors() == "dark" {
			return "Wechsle zu hellen Farben"
		} else {
			return "Wechsle zu dunklen Farben"
		}
	}
}
*/
