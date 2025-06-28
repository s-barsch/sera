//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/routes"
	"g.rg-s.com/sera/go/server"
	"g.rg-s.com/sera/go/viewer"
)

func main() {
	s, err := server.LoadServer()
	if err != nil {
		log.Fatal(err)
	}

	v, err := viewer.NewViewer(s.Logger, s.Store, s.Engine, s.Users, s.LoadSafe)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8013", routes.Router(v)))
}
