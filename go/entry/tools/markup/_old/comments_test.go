package parser

import (
	"testing"
)

var tests = []map[string]string{
	map[string]string{
		"sample":  "A line with a //comment.",
		"outcome": "A line with a ",
	},
	map[string]string{
		"sample":  "A multiline. //comment here.\nAnother line.",
		"outcome": "A multiline. \nAnother line.",
	},
	map[string]string{
		"sample":  "A line with a Note{here}.",
		"outcome": "A line with a Note.",
	},
	map[string]string{
		"sample":  "A line with a/* special comment */.",
		"outcome": "A line with a.",
	},
}

func TestRemoveHidden(t *testing.T) {
	for _, m := range tests {
		str := RemoveHidden(m["sample"])
		t.Logf("Result:  %v", str)
		t.Logf("Outcome: %v", m["outcome"])
		if str != m["outcome"] {
			t.Logf("strings not equal")
			t.Fail()
		}
	}
}
