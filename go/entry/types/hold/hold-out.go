package file 

import (
	"fmt"
	"os"
)

type Holds []*Hold

func (h *Hold) Title(lang string) string {
	title := h.Info.Title(lang)
	if title != "" {
		return title
	}
	return h.Acronym()
}

type Subnav struct {
	Label map[string]string
	Items []interface{}
}

func (h *Hold) IsSymlink() bool {
	return false
	fi, err := os.Stat(h.File.Path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

func (h *Hold) Subnav() *Subnav {
	if h.Section() == "graph" {
		if h.Depth() == 0 {
			return &Subnav{
				Label: map[string]string{
					"de": "Jahre",
					"en": "Years",
				},
				Items: toInterface(h.Holds.Reverse()),
				//Items: toInterface(h.Holds),
			}
		}
		if h.Depth() == 1 {
			if len(h.Holds) < 2 {
				return nil
			}
			return &Subnav{
				//Items: toInterface(h.Holds.Reverse()),
				Items: toInterface(h.Holds),
			}
		}
	}
	return &Subnav{
		Items: toInterface(h.Holds),
	}
}

func toInterface(h Holds) []interface{} {
	is := []interface{}{}
	for _, x := range h {
		is = append(is, x)
	}
	return is
}

func (h *Hold) Section() string {
	if h.Mother == nil {
		return h.File.Id
	}
	return h.Mother.Section()
}

func (h *Hold) IsEmpty() bool {
	if len(h.Els) == 0 && len(h.Holds) == 0 {
		return true
	}
	return false
}

func (h *Hold) Link(lang string) *chain {
	cs := h.Crumbs(lang)
	if len(cs) < 1 {
		return &chain{}
	}
	return cs[len(cs)-1]
}

func (h *Hold) Acronym() string {
	return EncodeAcronym(h.Date)
}

func (h *Hold) Crumbs(lang string) []*chain {
	path := ""
	chain := h.Chain(lang)
	for i, c := range chain {
		path += "/" + c.Slug
		if i == len(chain)-1 && i != 0 {
			c.Slug = path + "-" + h.Acronym()
			break
		}
		c.Slug = path + "/"
	}
	return chain
}

func (h *Hold) Id() string {
	return h.Date.Format(Timestamp)
}

func monthAnchor(path string) string {
	if len(path) > 3 {
		month := len(path) - 3
		return path[:month] + "#" + path[month+1:]
	}
	return path
}

func (h *Hold) Permalink(lang string) string {
	switch h.File.Base() {
	case "index":
		return "/index"
	case "graph":
		return "/graph"
	}
	switch h.Section() {
	case "graph":
		if h.Depth() == 2 {
			return monthAnchor(h.Path(lang))
		}
		if h.Depth() < 3 {
			return h.Path(lang)
		}
	case "extra":
		return fmt.Sprintf("/%v", h.Info.Slug(lang))
	default:
		if h.Depth() < 2 {
			return h.Path(lang)
		}
	}
	return fmt.Sprintf("%v-%v", h.Path(lang), EncodeAcronym(h.Date))
}

func (h *Hold) Depth() int {
	if h.Mother == nil {
		return 0
	}
	return 1 + h.Mother.Depth()
}

func (h *Hold) Name(lang string) string {
	if slug := h.Info.Slug(lang); slug != "" {
		return slug
	}
	return h.File.Id
}

type chain struct {
	Slug, Title string
}

func (h *Hold) Path(lang string) string {
	path := ""
	for _, c := range h.Chain(lang) {
		path += "/" + c.Slug
	}
	return path
}

func (h *Hold) Chain(lang string) []*chain {
	c := &chain{
		Slug:  h.Name(lang),
		Title: h.Info.Title(lang),
	}
	if h.Mother == nil {
		return []*chain{c}
	}
	return append(h.Mother.Chain(lang), c)
}

func (h *Hold) ImgLen() int {
	i := 0
	for _, e := range h.Els {
		_, ok := e.(*Image)
		if ok {
			i++
		}
	}
	return i
}
