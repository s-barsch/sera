package meta

import (
	"fmt"
	"sacer/go/entry"
	"sacer/go/server/users"
	"sacer/go/server/paths"
	"net/http"
)

type Meta struct {
	Auth    *users.Auth
	Options *Options

	Path    string
	Host    string

	Title   string
	Section string

	Nav   Nav
	Lang  string
	Langs Langs

	Desc   string
	Schema *Schema
}

func NewMeta(users *users.Users, r *http.Request) (*Meta, error) {
	auth, err := users.CheckAuth(r)
	if err != nil && err != http.ErrNoCookie {
		return nil, err
	}

	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		return nil, err
	}

	return &Meta{
		Auth:	 auth,
		Path:	 path,
		Lang:    Lang(r.Host),
		Options: GetOptions(r),
	}, nil
}

func (m *Meta) O() *Options {
	return m.Options
}

func GetAuth(r *http.Request) (*users.Auth, error) {
	a, ok := r.Context().Value("auth").(*users.Auth)
	if !ok {
		return nil, fmt.Errorf("head.GetAuth type assertion failed.")
	}
	return a, nil
}

func (m *Meta) Process(e entry.Entry) error {
	if m.Section == "" {
		return fmt.Errorf("section not set")
	}
	m.Desc = m.GetDesc(e)
	m.Langs = m.MakeLangs(e)
	m.Nav = m.MakeNav()

	return nil
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
