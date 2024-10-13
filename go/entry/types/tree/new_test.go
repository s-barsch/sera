package tree

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestReadTree(t *testing.T) {
	tree, err := ReadTree("/srv/rg-s/st/data/indecs/", nil)
	//s, err := ReadTree("./test/indecs", nil)
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
