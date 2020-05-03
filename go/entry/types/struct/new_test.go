package stru

import (
	"fmt"
	"github.com/kr/pretty"
	"testing"
)

func TestReadStruct(t *testing.T) {
	s, err := ReadStruct("/srv/rg-s/st/data/index/", nil)
	//s, err := ReadStruct("./test/index", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%# v", pretty.Formatter(s))
	printStruct(s)
}

func printStruct(str *Struct) {
	for _, s := range str.Structs {
		fmt.Println(s.Title("de"))
		fmt.Println(s.Perma("de"))
		printStruct(s)
	}
}
