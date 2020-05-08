package tmpl

import (
	"stferal/go/entry/types/tree"
)

type Subnav struct {
	Tree   *tree.Tree
	Active int64
	Lang   string
}

func NewSubnav(tree *tree.Tree, active int64, lang string) *Subnav {
	return &Subnav{
		Tree:   tree,
		Active: active,
		Lang:   lang,
	}
}

func (s *Subnav) T() *tree.Tree {
	return s.Tree
}

func (s *Subnav) L() string {
	return s.Lang
}

func (s *Subnav) NavTrees() tree.Trees {
	t := s.Tree
	if t.Section() == "graph" {
		if t.Level() == 0 {
			return t.Trees.Reverse()
		}
		// if only one month
		if len(t.Trees) < 2 {
			return nil
		}
	}
	return t.Trees
}

func (s *Subnav) IsYear() bool {
	return s.Tree.Level() == 0 && s.Tree.Section() == "graph"
}

var years = map[string]string{
	"de": "Jahre",
	"en": "Years",
}

func (s *Subnav) YearLabel(lang string) string {
	return years[lang]
}
