package paths

import (
	"github.com/kr/pretty"
	"reflect"
	"testing"
)

type testCase struct {
	path   string
	result *Path
}

var cases = []*testCase{
	&testCase{
		path: "/index/leben/ueberleben/oeffnungen-33ac2565",
		result: &Path{
			Page:    "index",
			Slugs:   []string{"leben", "ueberleben"},
			Slug:    "oeffnungen",
			Hash:    "33ac2565",
			Subdir:  "",
			Subpath: "",
		},
	},
	&testCase{
		path: "/index/kunst/innen-aussen-35e1fcdd",
		result: &Path{
			Page:    "index",
			Slugs:   []string{"kunst"},
			Slug:    "innen-aussen",
			Hash:    "35e1fcdd",
			Subdir:  "",
			Subpath: "",
		},
	},
	&testCase{
		path: "/index/kunst/form-34a1a15e",
		result: &Path{
			Page:    "index",
			Slugs:   []string{"kunst"},
			Slug:    "form",
			Hash:    "34a1a15e",
			Subdir:  "",
			Subpath: "",
		},
	},
	&testCase{
		path: "/graph/2020/03/09-36e55605/cache/200310_012140-1280.jpg",
		result: &Path{
			Page:    "graph",
			Slugs:   []string{"2020", "03"},
			Slug:    "09",
			Hash:    "36e55605",
			Subdir:  "cache",
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
