package server

import (
	"flag"
	"fmt"
	"log"
	"os"
	p "path/filepath"
	"sacer/go/entry/tools"
	"sacer/go/entry/tools/hyph"
	"sacer/go/server/tmpl"
	"sacer/go/server/users"
	"text/template"


	"github.com/rjeczalik/notify"
)

type Server struct {
	Paths *paths
	Flags *flags
	Log   *log.Logger

	Users *users.Users

	Trees   map[string]*DoubleTree
	Recents map[string]*DoubleEntries

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
	Info   bool
}

func NewServer() *Server {
	host := flag.String("host", "", "override host variable for testing")
	path := flag.String("path", ".", "set the root path of this app")
	all := flag.Bool("a", false, "sets all flags except mobile")
	debug := flag.Bool("debug", false, "log to stdout")
	local := flag.Bool("local", false, "enable local testing")
	reload := flag.Bool("reload", false, "reload files on every request")
	mobile := flag.Bool("mobile", false, "adjust polyfill path")
	info := flag.Bool("info", false, "display more video infos")

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
		Info: *info,
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

	err := LoadHyphPatterns(s.Paths.Root)
	if err != nil {
		return nil, err
	}

	u, err := users.LoadUsers()
	if err != nil {
		return nil, err
	}

	s.Users = u

	return s, s.Load()
}

func (s *Server) CloseUsers() error {
	err := s.Users.Close()
	if err != nil {
		return err
	}
	if s.Flags.Debug {
		log.Println("Closed user database.")
	}
	return nil
}

func newLogger(debug bool) *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func (s *Server) Debug(err error) {
	if s.Flags.Debug {
		s.Log.Println(err)
	}
}

func LoadHyphPatterns(root string) error {
	patterns := hyph.LangPatterns{}
	for lang, _ := range tools.Langs {
		path := fmt.Sprintf("%v/go/entry/tools/hyph/hyph-%v.dic", root, lang)
		p, err := hyph.LoadPattern(path)
		if err != nil {
			return err
		}
		patterns[lang] = p
	}
	hyph.SetLangPatterns(patterns)
	return nil
}
