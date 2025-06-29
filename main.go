//go:generate go/entry/types/generate

package main

import (
	"log"
	"net/http"

	"g.rg-s.com/sera/go/app"
	"g.rg-s.com/sera/go/server/flags"
)

func main() {
	a, err := app.Create(flags.Parse())
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8013", a.Router))
}
