package head

import (
	"stferal/pkg/el"
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
	// Image *el.Image
	// Location
}

func (h *Head) ElSchema() (*Schema, error) {
	date, err := el.DateSafe(h.El)
	if err != nil {
		return nil, err
	}

	typ := schemaType(h.El)

	perma, err := el.PermalinkSafe(h.El, h.Lang)
	if err != nil {
		return nil, err
	}

	url := h.AbsoluteURL(perma, h.Lang)

	mod, err := el.ModTimeSafe(h.El)
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
	if el.Type(e) == "image" {
		return "image"
	}
	return "article"
}
