package api

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io"
)

type User struct {
	Name string
	Mail string
	PaypalID string `json:"subscriptionID"`
}

func Main(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	err := json.NewDecoder(io.Reader(r.Body)).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = add(user)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%v", r)
	return
}


