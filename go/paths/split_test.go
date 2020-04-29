package paths

import (
	"github.com/kr/pretty"
	"testing"
)

var samples = []string{
	"/index/leben/ueberleben/oeffnungen-33ac2565",
	"/graph/2020/03/09-36e55605/cache/200310_012140-1280.jpg",
	"/graph/2020/03/14-3757ceb6/files/200116_235849.mp4",
}

func TestSplit(t *testing.T) {
	for _, path := range samples {
		p := Split(path)
		t.Logf("\n\n%v\n\n", path)
		t.Logf("%# v", pretty.Formatter(p))
	}
}
