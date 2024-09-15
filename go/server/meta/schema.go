package meta

import (
	//"g.rg-s.com/sferal/go/entry"
	"time"
)

type Schema struct {
	Title string
	Type  string // Article, Image, Audio, Video
	Url   string

	Date    time.Time
	ModDate time.Time

	Logo string
	Host string

	Description string
	// Image *entry.Image
	// Location
}

/*
func (m *Meta) ElSchema() (*Schema, error) {
	date, err := entry.DateSafe(h.El)
	if err != nil {
		return nil, err
	}

	typ := schemaType(h.El)

	perma, err := entry.PermalinkSafe(h.El, h.Lang)
	if err != nil {
		return nil, err
	}

	url := h.AbsoluteURL(perma, h.Lang)

	mod, err := entry.ModTimeSafe(h.El)
	if err != nil {
		return nil, err
	}

	desc := h.GetDesc()

	return &Schema{
		Title: h.Title,
		Type:  typ,
		Url:   url,

		Date:    date,
		ModDate: mod,

		Description: desc,

		Logo: h.AbsoluteURL("/static/img/pine.jpg", h.Lang),
		Host: h.Lang,
	}, nil
}

func schemaType(e interface{}) string {
	if entry.Type(e) == "image" {
		return "image"
	}
	return "article"
}
*/
