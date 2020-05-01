package head

import (
	/*
	"stferal/go/entry"
	"strings"
	*/
)

func (h *Head) GetDesc() string {
	return ""
	/*
	if h.Desc != "" {
		return h.Desc
	}

	d := entry.InfoSafe(h.El).Field("description", h.Lang)
	if d != "" {
		return d
	}

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
