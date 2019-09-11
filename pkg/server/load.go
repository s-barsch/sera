package server

import (
	"stferal/pkg/el"
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

	trees := map[string]*el.Hold{}

	for _, section := range sections {
		t, err := el.ReadHold(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}
		trees[section] = t
	}

	recents := map[string]el.Els{}
	recents["index"] = trees["index"].TraverseEls().Desc().NoHidden()
	recents["graph"] = trees["graph"].TraverseElsReverse()

	s.Trees = trees
	s.Recents = recents

	return nil
}
