package server

import (
	"stferal/pkg/entry"
	"time"
)

/*
import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)
*/

func (s *Server) Load() error {
	timeStart := time.Now()

	err := s.loadRender()
	if err != nil {
		return err

	}

	err = s.LoadData()
	if err != nil {
		return err
	}

	timeEnd := time.Now()
	timeDif := timeEnd.Sub(timeStart)

	if s.Flags.Debug {
		s.Log.Printf("Load: %v.\n", timeDif)
	}

	return nil
}

func (s *Server) LoadData() error {
	sections := []string{
		"index",
		"graph",
		"about",
		"extra",
	}

	trees := map[string]*entry.Hold{}

	for _, section := range sections {
		t, err := entry.ReadHold(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}
		trees[section] = t
	}

	recents := map[string]entry.Els{}
	recents["index"] = entry.PublicTree(trees["index"]).TraverseEls().Desc().Exclude()
	recents["graph"] = entry.PublicTree(trees["graph"]).TraverseElsReverse()

	s.Trees = trees
	s.Recents = recents

	return nil
}
