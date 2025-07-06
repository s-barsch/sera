package split

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

type testCase struct {
	path   string
	result *Split
}

var cases = []*testCase{
	{
		path: "/de/indecs/leben/ueberleben/oeffnungen-33ac2565",
		result: &Split{
			Raw:    "/de/indecs/leben/ueberleben/oeffnungen-33ac2565",
			Chain:  []string{"de", "indecs", "leben", "ueberleben"},
			Slug:   "oeffnungen",
			Hash:   "33ac2565",
			Folder: "",
			File:   nil,
		},
	},
	{
		path: "/en/cache/24-08/10-super-theory-3f8b02/img/cover-480.webp",
		result: &Split{
			Raw:    "/en/cache/24-08/10-super-theory-3f8b02/img/cover-480.webp",
			Chain:  []string{"en", "cache", "24-08"},
			Slug:   "10-super-theory",
			Hash:   "3f8b02",
			Folder: "img",
			File: &File{
				Name:   "cover.webp",
				Option: "480",
				Ext:    "webp",
			},
		},
	},
}

func TestSplit(t *testing.T) {
	for _, c := range cases {
		p := SplitPath(c.path)
		if !reflect.DeepEqual(p, c.result) {
			t.Errorf("Split failed. Want result:\n\n%# v", pretty.Formatter(c.result))
		}
		t.Logf("Sample path:\n\n%v\n\n", c.path)
		t.Logf("Split produced:\n\n%# v", pretty.Formatter(p))
	}
}

// TestSplitSlugHash tests the splitSlugHash function
func TestSplitSlugHash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantSlug string
		wantHash string
	}{
		{"Normal case", "lonely-3f397f82", "lonely", "3f397f82"},
		{"No hash", "just-a-slug", "just-a-slug", ""},
		{"Merged months", "11-12", "11-12", ""},
		{"All hash", "3f397f82", "", "3f397f82"},
		{"Empty string", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlug, gotHash := splitSlugHash(tt.input)
			if gotSlug != tt.wantSlug || gotHash != tt.wantHash {
				t.Errorf("splitSlugHash(%q) = (%q, %q), want (%q, %q)", tt.input, gotSlug, gotHash, tt.wantSlug, tt.wantHash)
			}
		})
	}
}
