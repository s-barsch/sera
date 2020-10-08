package head

import (
/*
	"sacer/go/entry"
	"strings"
*/
)

func (h *Head) GetDesc() string {
		if h.Desc != "" {
			return h.Desc
		}

		if h.Entry != nil {
			return h.Entry.Info().Field("description", h.Lang)
		}
		return ""
		/*
		d = h.MakeElDesc()
		if d != "" {
			return d
		}
		return d
		*/
}

func (h *Head) MakeElDesc() string {
	return ""
	/*
		desc := ""
		switch entry.Type(h.El) {
		case "image":
			desc = h.El.(*entry.Image).Info.Alt(h.Lang)
		case "text":
			desc = h.El.(*entry.Text).MetaDesc(h.Lang)
		}
		return strings.TrimSpace(desc)
	*/
}
