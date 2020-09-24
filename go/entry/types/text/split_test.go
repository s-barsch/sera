package text

import (
	"io/ioutil"
	"testing"
)

var testDir = "./test/"

func TestSplitTextFile(t *testing.T) {
	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, fi := range files {
		parts, err := splitTextFile(testDir + fi.Name())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", parts)
	}
}
