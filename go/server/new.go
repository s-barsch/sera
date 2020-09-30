package server

import (
	"flag"
	"github.com/rjeczalik/notify"
	"log"
	"os"
	p "path/filepath"
	"sacer/go/server/tmpl"
	"text/template"
)

type Server struct {
	Paths *paths
	Flags *flags
	Log   *log.Logger

	Trees   map[string]*SectionTree
	Recents map[string]*SectionEntries

	Templates *template.Template
	Vars      *tmpl.Vars

	Queue   chan int
	Watcher chan notify.EventInfo
	Quit    chan os.Signal
}

type paths struct {
	Root string
	Data string
}

type flags struct {
	Host   string
	Local  bool
	Debug  bool
	Reload bool
	Mobile bool
}

func NewServer() *Server {
	host := flag.String("host", "", "override host variable for testing")
	path := flag.String("path", ".", "set the root path of this app")
	all := flag.Bool("a", false, "sets all flags except mobile")
	debug := flag.Bool("debug", false, "log to stdout")
	local := flag.Bool("local", false, "enable local testing")
	reload := flag.Bool("reload", false, "reload files on every request")
	mobile := flag.Bool("mobile", false, "adjust polyfill path")

	flag.Parse()

	if *all {
		*debug = true
		*local = true
	}

	s := &Server{
		Log:   newLogger(*debug),
		Queue: make(chan int, 1),
	}

	s.Paths = &paths{
		Root: *path,
		// `Clean` is necessary to harmonize this path with later paths
		// that are processed by path/filepath functions.
		Data: p.Clean(*path + "/data"),
	}

	s.Flags = &flags{
		Host:   *host,
		Debug:  *debug,
		Local:  *local,
		Reload: *reload,
		Mobile: *mobile,
	}

	return s
}

func LoadServer() (*Server, error) {
	s := NewServer()

	if s.Flags.Debug {
		err := s.SetupWatcher()
		if err != nil {
			return nil, err
		}
	}

	return s, s.Load()
}

func newLogger(debug bool) *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func (s *Server) Debug(err error) {
	if s.Flags.Debug {
		s.Log.Println(err)
	}
}
