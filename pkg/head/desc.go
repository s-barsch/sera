package head

import (
	"st/pkg/el"
	"strings"
)

func (h *Head) GetDesc() string {
	if h.Desc != "" {
		return h.Desc
	}

	d := el.InfoSafe(h.El).Field("description", h.Lang)
	if d != "" {
		return d
	}

	d = h.MakeElDesc()
	if d != "" {
		return d
	}
	return d
}

func (h *Head) MakeElDesc() string {
	desc := ""
	switch el.Type(h.El) {
	case "image":
		desc = h.El.(*el.Image).Info.Alt(h.Lang)
	case "text":
		desc = h.El.(*el.Text).MetaDesc(h.Lang)
	}
	return strings.TrimSpace(desc)
}
