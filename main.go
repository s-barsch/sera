//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/routes"
	"g.rg-s.com/sera/go/server"
)

func main() {

	s, err := server.LoadServer()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", routes.Router(s))
	log.Fatal(http.ListenAndServe(":8013", nil))
}
