package server

import (
	"log"
	"text/template"

	"g.rg-s.com/sacer/go/requests/tmpl"
	"g.rg-s.com/sacer/go/server/flags"
	"g.rg-s.com/sacer/go/server/store"

	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	config config

	store  *store.Store
	engine *Engine

	trigger chan struct{}
}

type config struct {
	Paths tmpl.Paths
	Flags Flags
}

type Flags struct {
	Debug bool
	Local bool
	Info  bool
}

type Engine struct {
	*template.Template
	Vars *tmpl.Vars
}

func New(logger *logrus.Logger, flags flags.Flags) (*Server, error) {
	log.SetFlags(log.LstdFlags)
	return &Server{
		logger: logger,
	}, nil
}

/*
func (s *Server) CloseUsers() error {
	err := s.Users.Close()
	if err != nil {
		return err
	}
	if s.Debug {
		log.Println("Closed user database.")
	}
	return nil
}
*/
