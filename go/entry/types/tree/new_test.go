package tree

import (
	"fmt"
	"github.com/kr/pretty"
	"testing"
)

func TestReadTree(t *testing.T) {
	s, err := ReadTree("/srv/rg-s/st/data/index/", nil)
	//s, err := ReadTree("./test/index", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%# v", pretty.Formatter(s))
	printTree(s)
}

func printTree(str *Tree) {
	for _, s := range str.Trees {
		fmt.Println(s.Title("de"))
		fmt.Println(s.Perma("de"))
		printTree(s)
	}
}
