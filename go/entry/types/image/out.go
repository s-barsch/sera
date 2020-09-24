package image

import (
	"fmt"
	p "path/filepath"
	"strconv"
)

func (i *Image) Location(size string) string {
	return p.Join(i.file.Dir(), "cache", size, i.file.Name())
}

func (i *Image) ImagePath(size int, lang string) string {
	if i.parent.Type() == "set" {
		return fmt.Sprintf("%v/cache/%v", i.parent.Perma(lang), i.ImageName(size))
	}
	return fmt.Sprintf("%v/cache/%v", i.Perma(lang), i.ImageName(size))
}

func (i *Image) ImageName(size int) string {
	return fmt.Sprintf("%v-%v%v", i.file.NameNoExt(), size, i.file.Ext())
}

func (i *Image) SrcSet(size int, lang string) string {
	return fmt.Sprintf("%v %vw", i.ImagePath(size, lang), i.Width(size))
}

/*
func (i *Image) Permalink(lang string) string {
	if i.File.Section() == "index" {
		return fmt.Sprintf("%v#%v", i.File.Hold.Permalink(lang), Normalize(i.Title(lang)))
	}
	if i.Info.Title(lang) == "" {
		return fmt.Sprintf("%v/%v", i.File.Hold.Path(lang), ToB16(i.Date))
	}
	return fmt.Sprintf("%v/%v-%v", i.File.Hold.Path(lang), i.Info.Slug(lang), i.Acronym())
}
*/

/*
func (i *Image) ImagePath(size int, lang string) string {
	if i.File.Hold.Info["read"] != "false" {
		return fmt.Sprintf("%v/cache/%v", i.File.Hold.Permalink(lang), i.ImageName(size))
	}
	return fmt.Sprintf("%v/cache/%v", i.Permalink(lang), i.ImageName(size))
}


*/

// dim related

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
	case "1.4":
		return "a4w"
	case "0.7":
		return "32h"
		//return "a4h"
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
