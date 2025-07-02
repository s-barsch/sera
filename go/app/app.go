package app

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/router"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/flags"
	"g.rg-s.com/sera/go/server/users"
	"g.rg-s.com/sera/go/viewer"
)

type App struct {
	Server *server.Server
	Viewer *viewer.Viewer
	Router http.Handler
}

func Create(flags flags.Flags) (App, error) {
	users, err := users.LoadUsers()
	if err != nil {
		return App{}, err
	}

	s, err := server.LoadServer(flags)
	if err != nil {
		log.Fatal(err)
	}

	v, err := viewer.NewViewer(s.Logger, s.Store, s.Engine, users, s.LoadSafe)
	if err != nil {
		log.Fatal(err)
	}

	return App{
		Viewer: v,
		Router: router.New(v),
	}, nil
}
