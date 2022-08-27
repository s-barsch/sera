package auth

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

	switch r.URL.Path {
	case "/api/login":
		Login(w, r)
	case "/api/subscribe":
		Subscribe(w, r)
	case "/api/register":
		Register(w, r)
	}

	//fmt.Printf("%v", r)
	return
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
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
}

func Register(w http.ResponseWriter, r *http.Request) {
	user := &User{}

	user.Name = r.FormValue("name")
	user.Mail = r.FormValue("mail")

	err := add(user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	mail := r.FormValue("mail")
	user, err := lookup(mail)
	if err != nil {
		fmt.Println(err)
	}

	err = send(user.Mail, "hello my friend")
	if err != nil {
		fmt.Println(err)
	}
	println("apperantly i sent")
}
