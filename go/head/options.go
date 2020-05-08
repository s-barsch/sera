package head

import (
	"net/http"
)

func GetOptions(r *http.Request) map[string]bool {
	return map[string]bool{
		"large": LargeType(r),
		"dark":  DarkColors(r),
	}
}

func LargeType(r *http.Request) bool {
	c, err := r.Cookie("type")
	if err != nil {
		return true
	}
	return c.Value == "large"
}

func DarkColors(r *http.Request) bool {
	c, err := r.Cookie("colors")
	if err != nil {
		return false
	}
	return c.Value == "dark"
}

func LogMode(r *http.Request) bool {
	c, err := r.Cookie("nolog")
	if err != nil {
		return false
	}
	return c.Value == "true"
}

func (h *Head) Dark() bool {
	if h.Options == nil {
		return false
	}
	return h.Options["dark"]
}

func (h *Head) Large() bool {
	if h.Options == nil {
		return false
	}
	return h.Options["large"]
}

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

func (h *Head) SwitchColorsTitle(lang string) string {
	switch lang {
	case "en":
		if h.Dark() {
			return "Switch to light colors"
		} else {
			return "Switch to dark colors"
		}
	default:
		if h.Dark() {
			return "Wechsle zu hellen Farben"
		} else {
			return "Wechsle zu dunklen Farben"
		}
	}
}

func (h *Head) SwitchTypeLink(lang string) string {
	if h.Large() {
		return "/opt/type/small"
	} else {
		return "/opt/type/large"
	}
}

func (h *Head) SwitchColorsLink(lang string) string {
	if h.Dark() {
		return "/opt/colors/light"
	} else {
		return "/opt/colors/dark"
	}
}
