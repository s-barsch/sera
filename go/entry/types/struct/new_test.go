package stru

import (
	"github.com/kr/pretty"
	"testing"
)

func TestReadStruct(t *testing.T) {
	s, err := ReadStruct("/srv/rg-s/st/data/index", nil)
	//s, err := ReadStruct("./test/index", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%# v", pretty.Formatter(s))
}
