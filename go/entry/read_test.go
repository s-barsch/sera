package entry

import (
	"github.com/kr/pretty"
	"testing"
)

func TestReadStructure(t *testing.T) {
	s, err := ReadStructure("./test/index", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%# v", pretty.Formatter(s))
}
