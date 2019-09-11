package main

import (
	"net/http"
	"st/pkg/server"
)

func main() {
	s := server.New()

	err := s.Load()
	if err != nil {
		s.Log.Println(err)
	}

	http.Handle("/", routes(s))
	http.ListenAndServe(":8013", nil)
}
