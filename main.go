//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/server/flags"
	"g.rg-s.com/sera/go/server/watcher"
	"github.com/sirupsen/logrus"
)

func main() {
	s, err := initServer()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8013", newRouter(s)))
}

func initServer() (*server.Server, error) {
	flags, err := flags.Parse()
	if err != nil {
		return nil, err
	}

	logger := logrus.New()

	s, err := server.New(logger, flags)
	if err != nil {
		return nil, err
	}

	err = watcher.Init(logger, s.Paths, s.Trigger)
	if err != nil {
		return nil, err
	}

	return s, nil
}
