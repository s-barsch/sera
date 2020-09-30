//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"
	"sacer/go/routes"
	"sacer/go/server"
)

func main() {

	s, err := server.LoadServer()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", routes.Router(s))
	http.ListenAndServe(":8013", nil)
}
