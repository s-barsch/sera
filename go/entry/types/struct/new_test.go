package entry

import (
	"github.com/kr/pretty"
	"testing"
)

func TestReadStructure(t *testing.T) {
	s, err := ReadStruct("./test/index", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%# v", pretty.Formatter(s))
}
