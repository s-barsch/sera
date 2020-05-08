package server

import (
	"stferal/go/server/tmpl"
	"time"
)

func (s *Server) Load() error {
	tStart := time.Now()

	err := s.LoadTemplates()
	if err != nil {
		return err

	}

	err = s.LoadTrees()
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

// templates

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
