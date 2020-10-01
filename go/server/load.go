package server

import (
	"fmt"
	"sacer/go/server/tmpl"
	"time"
)

var sections = []string{
	"index",
	"graph",
	"about",
	"extra",
	"kine",
}

func (s *Server) LoadSafe() error {
	select {
	case s.Queue <- 1:
		err := s.Load()
		<-s.Queue
		return err
	default:
		return fmt.Errorf("Load is blocked.")
	}
}

func (s *Server) Load() error {
	tStart := time.Now()

	err := s.ReadTrees()
	if err != nil {
		return err
	}

	err = s.processAllTexts()
	if err != nil {
		return err
	}

	s.makeLinks()

	err = s.LoadTemplates()
	if err != nil {
		return err

	}

	tEnd := time.Now()
	tDif := tEnd.Sub(tStart)

	if s.Flags.Debug {
		s.Log.Printf("Load: %v.\n", tDif)
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

	s.Templates = ts
	s.Vars = vars

	return nil
}
