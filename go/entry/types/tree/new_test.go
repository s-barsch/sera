package tree

import (
	"fmt"
	"github.com/kr/pretty"
	"testing"
)

func TestReadTree(t *testing.T) {
	tree, err := ReadTree("/srv/rg-s/st/data/index/", nil)
	//s, err := ReadTree("./test/index", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%# v", pretty.Formatter(tree))
	printTree(tree)
}

func printTree(tree *Tree) {
	for _, t := range tree.Trees {
		fmt.Println(t.Section())
		fmt.Println(t.Title("de"))
		fmt.Println(t.Perma("de"))
		// recursive function call
		printTree(t)
	}
}
