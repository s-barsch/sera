//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	router "g.rg-s.com/sacer/go/routes"
	"g.rg-s.com/sacer/go/server"
	"g.rg-s.com/sacer/go/server/flags"
	"g.rg-s.com/sacer/go/viewer"
	"github.com/sirupsen/logrus"
)

func main() {
	v, err := viewer.New(logrus.New(), nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8013", router.New(v)))
}

func initServer() (*server.Server, error) {
	flags := flags.Parse()

	logger := logrus.New()

	s, err := server.New(logger, flags)
	if err != nil {
		return nil, err
	}

	/*
		err = watcher.Init(logger, s.Paths, s.Trigger)
		if err != nil {
			return nil, err
		}
	*/

	return s, nil
}
