package tmpl

import (
	"sacer/go/entry/types/tree"
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
		/*
			// if only one month
			if len(t.Trees) < 2 {
				return nil
			}
		*/
	}
	if t.Section() == "kine" {
		return t.Trees.Reverse()
		/*
			if t.Level() == 1 {
				return t.Trees.Reverse()
			}
		*/
	}
	return t.Trees
}

func (s *Subnav) IsYear() bool {
	return s.Tree.Level() == 0 && s.Tree.Section() == "graph"
}

func (s *Subnav) IsDay() bool {
	section := s.Tree.Section()
	if section != "graph" && section != "kine" {
		return false
	}
	if s.Tree.Level() != 2 {
		return false
	}
	if s.Active == 0 {
		return false
	}
	return s.Tree.Id() != s.Active
}

var years = map[string]string{
	"de": "Jahre",
	"en": "Years",
}

func (s *Subnav) YearLabel(lang string) string {
	return years[lang]
}
