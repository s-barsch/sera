package paths

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

type testCase struct {
	path   string
	result *Path
}

var cases = []*testCase{
	{
		path: "/de/indecs/leben/ueberleben/oeffnungen-33ac2565",
		result: &Path{
			Path:   "/de/indecs/leben/ueberleben/oeffnungen-33ac2565",
			Chain:  []string{"de", "indecs", "leben", "ueberleben"},
			Slug:   "oeffnungen",
			Hash:   "33ac2565",
			Folder: "",
			File:   &File{},
		},
	},
	{
		path: "/en/cache/24-08/10-well-i-call-it-art-theory-3f358b02/files/sizes/240810_094201-1080.mp4",
		result: &Path{
			Path:   "/en/cache/24-08/10-well-i-call-it-art-theory-3f358b02/files/sizes/240810_094201-1080.mp4",
			Chain:  []string{"en", "cache", "24-08"},
			Slug:   "10-well-i-call-it-art-theory",
			Hash:   "3f358b02",
			Folder: "files",
			File: &File{
				Name:   "sizes/240810_094201.mp4",
				Option: "1080",
				Ext:    "mp4",
			},
		},
	},
	{
		path: "/en/cache/24-08/10-well-i-call-it-art-theory-3f358b02/img/cover-480.webp",
		result: &Path{
			Path:   "/en/cache/24-08/10-well-i-call-it-art-theory-3f358b02/img/cover-480.webp",
			Chain:  []string{"en", "cache", "24-08"},
			Slug:   "10-well-i-call-it-art-theory",
			Hash:   "3f358b02",
			Folder: "img",
			File: &File{
				Name:   "cover.webp",
				Option: "480",
				Ext:    "webp",
			},
		},
	},
	{
		path: "/en/img/24-08/10-well-i-call-it-art-theory-3f358b02/img/cover-480.webp",
		result: &Path{
			Path:   "/en/img/24-08/10-well-i-call-it-art-theory-3f358b02/img/cover-480.webp",
			Chain:  []string{"en", "img", "24-08"},
			Slug:   "10-well-i-call-it-art-theory",
			Hash:   "3f358b02",
			Folder: "img",
			File: &File{
				Name:   "cover.webp",
				Option: "480",
				Ext:    "webp",
			},
		},
	},
	/*
		&testCase{
			path: "/indecs/leben/ueberleben/oeffnungen-33ac2565",
			result: &Path{
				Page:    "indecs",
				Slugs:   []string{"leben", "ueberleben"},
				Slug:    "oeffnungen",
				Hash:    "33ac2565",
				Subdir:  "",
				Subpath: "",
			},
		},
		&testCase{
			path: "/indecs/kunst/innen-aussen-35e1fcdd",
			result: &Path{
				Page:    "indecs",
				Slugs:   []string{"kunst"},
				Slug:    "innen-aussen",
				Hash:    "35e1fcdd",
				Subdir:  "",
				Subpath: "",
			},
		},
		&testCase{
			path: "/indecs/kunst/form-34a1a15e",
			result: &Path{
				Page:    "indecs",
				Slugs:   []string{"kunst"},
				Slug:    "form",
				Hash:    "34a1a15e",
				Subdir:  "",
				Subpath: "",
			},
		},
		&testCase{
			path: "/graph/2020/03/09-36e55605/img/200310_012140-1280.jpg",
			result: &Path{
				Page:    "graph",
				Slugs:   []string{"2020", "03"},
				Slug:    "09",
				Hash:    "36e55605",
				Subdir:  "img",
				Subpath: "200310_012140-1280.jpg",
			},
		},
		&testCase{
			path: "/graph/2020/03/14-3757ceb6/files/200116_235849.mp4",
			result: &Path{
				Page:    "graph",
				Slugs:   []string{"2020", "03"},
				Slug:    "14",
				Hash:    "3757ceb6",
				Subdir:  "files",
				Subpath: "200116_235849.mp4",
			},
		},
	*/
}

func TestSplit(t *testing.T) {
	for _, c := range cases {
		p := Split(c.path)
		if !reflect.DeepEqual(p, c.result) {
			t.Errorf("Split failed. Want result:\n\n%# v", pretty.Formatter(c.result))
		}
		t.Logf("Sample path:\n\n%v\n\n", c.path)
		t.Logf("Split produced:\n\n%# v", pretty.Formatter(p))
	}
}
