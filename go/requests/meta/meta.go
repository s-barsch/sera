package meta

import (
	"fmt"
	"net/http"

	"g.rg-s.com/sacer/go/entry"
	"g.rg-s.com/sacer/go/requests/split"
)

type Meta struct {
	*Options

	Path string
	Host string

	Split *split.Split

	Title   string
	section string

	Nav   Nav
	Lang  string
	Langs Langs

	Desc string
	*Schema
}

func (m *Meta) Section() string {
	return m.section
}

func NewMeta(r *http.Request) (*Meta, error) {
	path, err := split.Sanitize(r.URL.Path)
	if err != nil {
		return nil, err
	}

	return &Meta{
		Path:    path,
		Split:   split.SplitPath(path),
		Lang:    Lang(path),
		Host:    r.Host,
		Options: GetOptions(r),
	}, nil
}

func (m *Meta) O() *Options {
	return m.Options
}

func (m *Meta) SetHreflang(e entry.Entry) {
	m.Langs = MakeHreflangs(m.HostAddress(), e)
}

func (m *Meta) SetSection(section string) {
	m.section = section
	m.SetNav(section)
}

func SiteName(lang string) string {
	return "Sacer Feral"
}

func (m *Meta) PageTitle() string {
	name := SiteName(m.Lang)
	if m.Title == "" {
		return name
	}
	return fmt.Sprintf("%v - %v", m.Title, name)
}

func (m *Meta) PageURL() string {
	if l := m.Langs.Hreflang(m.Lang); l != nil {
		return l.Href
	}
	return ""
}

func (m *Meta) DontIndex() bool {
	switch m.Path {
	case "/impressum", "/legal", "/privacy", "/datenschutz":
		return true
	}
	return false
}
