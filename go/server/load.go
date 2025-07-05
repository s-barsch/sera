package server

import (
	"fmt"
	"log"
	"time"

	"g.rg-s.com/sera/go/server/tmpl"
)

func (s *Server) LoadSafe() error {
	select {
	case s.Queue <- 1:
		err := s.Load()
		<-s.Queue
		return err
	default:
		return fmt.Errorf("load is blocked")
	}
}

func (s *Server) Load() error {
	tStart := time.Now()

	/*
		err := s.ReadTrees()
		if err != nil {
			return err
		}
	*/

	err := s.LoadTemplates()
	if err != nil {
		return err

	}

	tEnd := time.Now()
	tDif := tEnd.Sub(tStart)

	if s.Flags.Debug {
		log.Printf("Load: %v.\n", tDif)
	}

	return nil
}

func (s *Server) LoadTemplates() error {
	vars, err := tmpl.LoadVars(s.Paths.Root)
	if err != nil {
		return err
	}

	ts, err := tmpl.LoadTemplates(s.Paths.Root, s.Funcs())
	if err != nil {
		return err
	}

	_ = ts
	_ = vars

	/*
		s.Engine.Templates = ts
		s.Engine.Vars = vars
	*/

	return nil
}
