//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	"g.sacerb.com/sacer/go/routes"
	"g.sacerb.com/sacer/go/server"
)

func main() {

	s, err := server.LoadServer()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", routes.Router(s))
	http.ListenAndServe(":8013", nil)
}
