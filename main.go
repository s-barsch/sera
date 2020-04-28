package main

import (
	"net/http"
	"stferal/go/server"
	"stferal/go/routes"
)

func main() {
	s := server.New()

	err := s.Load()
	if err != nil {
		s.Log.Println(err)
	}

	http.Handle("/", routes.Router(s))
	http.ListenAndServe(":8013", nil)
}
