//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/routes"
	"g.rg-s.com/sera/go/server"
	s "g.rg-s.com/sera/go/server"
)

func main() {
	store, err := server.LoadServer()
	if err != nil {
		log.Fatal(err)
	}
	s.Srv = store

	http.Handle("/", routes.Router(store))
	log.Fatal(http.ListenAndServe(":8013", nil))
}
