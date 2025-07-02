package server

import (
	"log"
	"os"
	"text/template"

	"g.rg-s.com/sera/go/server/flags"
	"g.rg-s.com/sera/go/server/tmpl"

	"github.com/rjeczalik/notify"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Flags
	Paths  tmpl.Paths
	Logger *logrus.Logger

	Store  *Store
	Engine *Engine

	Queue   chan int
	Watcher chan notify.EventInfo
	Quit    chan os.Signal
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

type Store struct {
	Trees   map[string]*DoubleTree
	Recents map[string]*DoubleEntries
}

func NewServer(flags flags.Flags) *Server {

	s := &Server{
		Flags: Flags{
			Debug: flags.Debug,
			Local: flags.Local,
			Info:  flags.Info,
		},
		Queue: make(chan int, 1),
	}
	return s
}

func LoadServer(flags flags.Flags) (*Server, error) {
	s := NewServer(flags)

	log.SetFlags(log.LstdFlags)

	if flags.Debug {
		err := s.SetupWatcher()
		if err != nil {
			return nil, err
		}
	}

	return s, s.Load()
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
