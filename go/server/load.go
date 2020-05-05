package server

import (
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
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

	trees := map[string]*tree.Tree{}
	recents := map[string]entry.Entries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}
		//trees[section+"-private"] = t
		//trees[section] = t.Public()
		trees[section] = t

		recents[section] = trees[section].TraverseEntries()
		/*
		if section == "graph" {
			recents[section+"-private"] = trees[section+"-private"].TraverseElsReverse()
		} else {
			recents[section+"-private"] = trees[section+"-private"].TraverseEls().Desc().Exclude()
		}

		recents[section] = recents[section+"-private"].Public()
		*/
	}

	s.Trees = trees
	s.Recents = recents

	return nil
}
