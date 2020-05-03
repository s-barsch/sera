package server

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"strings"
	"text/template"
)

type Server struct {
	Paths *paths
	Flags *flags
	Log   *log.Logger

	Trees   map[string]*tree.Tree
	Recents map[string]entry.Entries

	Templates *template.Template
	Vars      vars
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

func New() *Server {
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
		*reload = true
	}

	s := &Server{}

	p, err := readConfFile()
	if err == nil {
		*path = p
	}

	s.Paths = &paths{
		Root: *path,
		Data: *path + "/data",
	}

	s.Flags = &flags{
		Host:   *host,
		Debug:  *debug,
		Local:  *local,
		Reload: *reload,
		Mobile: *mobile,
	}

	s.Log = newLogger(s.Flags.Debug)

	return s
}

func newLogger(debug bool) *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

func (s *Server) Debug(err error) {
	if s.Flags.Debug {
		s.Log.Println(err)
	}
}

func readConfFile() (string, error) {
	b, err := ioutil.ReadFile("/etc/stferal.conf")
	return strings.TrimSpace(string(b)), err
}

/*
import (
	"fmt"
	"log"
	"net/http"
	"st/el"
	"strings"
	"text/template"
	"time"
)

type paths struct {
	root string
	data string
	app  string
}

type server struct {
	paths paths

	indexGraph map[string]string
	logo       string

	bundleModTime time.Time

	flags *flags

	indexRecent entry.Els
	graphRecent entry.Els

	indexTree *entry.Hold
	graphTree *entry.Hold
	aboutTree *entry.Hold
	extraTree *entry.Hold

	vars map[string]string

	tmpls *template.Template
}

var srv server
*/
