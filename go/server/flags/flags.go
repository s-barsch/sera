package flags

import (
	"flag"
	"path/filepath"
)

type Flags struct {
	Host     string
	RootPath string
	DataPath string
	Local    bool
	Debug    bool
	Reload   bool
	Mobile   bool
	Info     bool
}

func Parse() Flags {
	host := flag.String("host", "", "override host variable for testing")
	path := flag.String("path", ".", "set the root path of this app")
	all := flag.Bool("a", false, "sets debug and local to true")
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

	return Flags{
		Host:     *host,
		RootPath: *path,
		DataPath: filepath.Clean(*path + "/data"),
		Debug:    *debug,
		Local:    *local,
		Reload:   *reload,
		Mobile:   *mobile,
		Info:     *info,
	}
}
