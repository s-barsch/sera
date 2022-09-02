package meta

import (
	"sacer/go/entry"
/*
	"strings"
*/
)

func (m *Meta) GetDesc(e entry.Entry) string {
		if m.Desc != "" {
			return m.Desc
		}

		if e != nil {
			return e.Info().Field("description", m.Lang)
		}
		return ""
		/*
		d = m.MakeElDesc()
		if d != "" {
			return d
		}
		return d
		*/
}

func (m *Meta) MakeElDesc() string {
	return ""
	/*
		desc := ""
		switch entry.Type(m.El) {
		case "image":
			desc = m.El.(*entry.Image).Info.Alt(m.Lang)
		case "text":
			desc = m.El.(*entry.Text).MetaDesc(m.Lang)
		}
		return strings.TrimSpace(desc)
	*/
}
