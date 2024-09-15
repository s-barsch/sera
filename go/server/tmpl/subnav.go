package tmpl

import (
	"g.rg-s.com/sferal/go/entry"
	"g.rg-s.com/sferal/go/entry/types/tree"
)

type Subnav struct {
	Tree   *tree.Tree
	Active entry.Entry
	Lang   string
}

func NewSubnav(tree *tree.Tree, active entry.Entry, lang string) *Subnav {
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

func (s *Subnav) ActiveId() int64 {
	if s.Active == nil {
		return 0
	}
	return s.Active.Id()
}

func (s *Subnav) NavTrees() tree.Trees {
	if s.Tree == nil {
		panic("subnav Tree shouldnt be empty")
	}
	t := s.Tree
	if t.Section() == "graph" || t.Section() == "cache" {
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
	if t.Section() == "cache" {
		/*
			return t.Trees.Reverse()
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
	if section != "graph" && section != "cache" {
		return false
	}
	if s.Tree.Level() != 2 {
		return false
	}
	if s.Active == nil {
		return false
	}
	return s.Tree.Id() != s.Active.Id()
}

var years = map[string]string{
	"de": "Jahre",
	"en": "Years",
}

func (s *Subnav) YearLabel(lang string) string {
	return years[lang]
}
