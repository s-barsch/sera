package head

import (
	"net/http"
)

type Options struct {
	Colors string 
	Size   string
}

func GetOptions(r *http.Request) *Options {
	return &Options {
		Colors: GetColors(r),
		Size:   GetSize(r),
	}
}

func GetSize(r *http.Request) string {
	c, err := r.Cookie("size")
	if err != nil {
		return "medium"
	}
	switch c.Value {
	case "small", "medium", "large":
		return c.Value
	}
	return "medium"
}

func GetColors(r *http.Request) string {
	c, err := r.Cookie("colors")
	if err != nil {
	}
	switch c.Value {
	case "dark", "light":
		return c.Value
	}
	return "light"
}

func (h *Head) Colors() string {
	if h.Options == nil {
		return "light"
	}
	return h.Options.Colors
}

func (h *Head) Size() string {
	if h.Options == nil {
		return "medium"
	}
	return h.Options.Size
}

/*
func (h *Head) SwitchTypeTitle(lang string) string {
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
*/

func (h *Head) SwitchColorsTitle(lang string) string {
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

func (h *Head) SwitchTypeLink(lang string) string {
	switch h.Size() {
	case "medium":
		return "/opt/size/large"
	case "large":
		return "/opt/size/small"
	}
	return "/opt/size/medium"
}

func (h *Head) SwitchColorsLink(lang string) string {
	if h.Colors() == "light" {
		return "/opt/colors/light"
	} else {
		return "/opt/colors/dark"
	}
}
