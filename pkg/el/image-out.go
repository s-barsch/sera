package el

import (
	"fmt"
	"path/filepath"
	"strconv"
)

func (i *Image) Acronym() string {
	return ToB16(i.Date)
}

func (i *Image) AcronymShort() string {
	return shortenAcronym(i.Acronym())
}

func (i *Image) Title(lang string) string {
	t := i.Info.Title(lang)
	if t != "" {
		return t
	}
	return i.AcronymShort()
}

func (i *Image) Permalink(lang string) string {
	if i.Info.Title(lang) == "" {
		return fmt.Sprintf("%v/%v", i.File.Hold.Path(lang), ToB16(i.Date))
	}
	return fmt.Sprintf("%v/%v-%v", i.File.Hold.Path(lang), i.Info.Slug(lang), i.Acronym())
}

func (i *Image) Ratio() float64 {
	w := i.Dims.width
	h := i.Dims.height
	return float64(w) / float64(h)
}

func (i *Image) RatioCode() string {
	switch x := fmt.Sprintf("%.1f", i.Ratio()); x {
	case "0.8":
		return "43h"
	case "1.3":
		return "43w"
	case "0.7":
		return "32h"
	case "1.5":
		return "32w"
	case "0.6":
		return "169h"
	case "1.8":
		return "169w"
	default:
		return x
	}
}

func (i *Image) PlaceholderPath() string {
	return fmt.Sprintf("/static/img/placeholder/%v.jpg", i.RatioCode())
}

func (i *Image) Orientation() string {
	if i.Dims.width >= i.Dims.height {
		return "landscape"
	}
	return "portrait"
}

func (i *Image) Height(size int) string {
	w := i.Dims.width
	h := i.Dims.height
	if w <= h {
		return strconv.Itoa(size)
	}
	return strconv.Itoa(int(float64(size) / i.Ratio()))
}

func (i *Image) Width(size int) string {
	w := i.Dims.width
	h := i.Dims.height
	if w >= h {
		return strconv.Itoa(size)
	}
	return strconv.Itoa(int(i.Ratio() * float64(size)))
}

func (i *Image) ImageAbs(size string) string {
	return filepath.Join(cacheFolder(i.File.Path), size, filepath.Base(i.File.Path))
}

func (i *Image) SrcSet(size int, lang string) string {
	return fmt.Sprintf("%v %vw", i.ImagePath(size, lang), i.Width(size))
}

func (i *Image) ImagePath(size int, lang string) string {
	if i.File.Hold.Info["read"] != "false" {
		return fmt.Sprintf("%v/cache/%v", i.File.Hold.Permalink(lang), i.ImageName(size))
	}
	return fmt.Sprintf("%v/cache/%v", i.Permalink(lang), i.ImageName(size))
}

func (i *Image) ImageName(size int) string {
	return fmt.Sprintf("%v-%v%v", i.File.BaseNoExt(), size, i.File.Ext())
}
