package meta

import (
	"fmt"
	"net/http"

	"g.rg-s.com/sera/go/entry"
	"g.rg-s.com/sera/go/server/paths"
	usr "g.rg-s.com/sera/go/server/users"
)

type Meta struct {
	*usr.Auth
	*Options

	Path string
	Host string

	Split *paths.Split

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

func NewMeta(users *usr.Users, w http.ResponseWriter, r *http.Request) (*Meta, error) {
	auth, err := users.CheckAuth(r)
	if err != nil && err != http.ErrNoCookie {
		usr.DeleteSessionCookie(w)
		return nil, err
	}

	path, err := paths.Sanitize(r.URL.Path)
	if err != nil {
		return nil, err
	}

	return &Meta{
		Auth:    auth,
		Path:    path,
		Split:   paths.SplitPath(path),
		Lang:    Lang(path),
		Host:    r.Host,
		Options: GetOptions(r),
	}, nil
}

func (m *Meta) O() *Options {
	return m.Options
}

func GetAuth(r *http.Request) (*usr.Auth, error) {
	a, ok := r.Context().Value("auth").(*usr.Auth)
	if !ok {
		return nil, fmt.Errorf("head.GetAuth type assertion failed")
	}
	return a, nil
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
