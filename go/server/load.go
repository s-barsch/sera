package server

import (
	"log"
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

func (s *Server) Reload() {
	select {
	case s.Queue <- 1:
		log.Println("Load started.")
	default:
		log.Println("Queue full.")
		return
	}
	if len(s.Queue) <= 1 {
		go s.runLoad()
	}
}

func (s *Server) runLoad() {
	err := s.Load()
	if err != nil {
		log.Println(err)
	}

	<-s.Queue

	if len(s.Queue) > 0 {
		go s.runLoad()
	}
}

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

	err = s.processAllTexts()
	if err != nil {
		return err
	}

	s.makeLinks()

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
