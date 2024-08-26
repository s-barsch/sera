package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"g.sacerb.com/sacer/go/server"
	"g.sacerb.com/sacer/go/server/users"
)

func Subscribe(s *server.Server, w http.ResponseWriter, r *http.Request) {
	user := &users.User{}

	err := json.NewDecoder(io.Reader(r.Body)).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = s.Users.AddUser(user)
	if err != nil {
		fmt.Println(err)
	}
}

func Register(s *server.Server, w http.ResponseWriter, r *http.Request) {
	user := &users.User{}

	user.Name = r.FormValue("name")
	user.Mail = r.FormValue("mail")

	err := s.Users.AddUser(user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
}

func RequestLogin(s *server.Server, w http.ResponseWriter, r *http.Request) {
	mail := r.FormValue("mail")
	user, err := s.Users.LookupUser(mail)
	if err != nil {
		log.Println(err)
		return
	}

	key, err := users.GenerateLoginKey()
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Users.StoreVerify(mail, key)
	if err != nil {
		log.Println(err)
		return
	}

	err = send(user.Mail, fmt.Sprintf(loginTmpl, users.EncodeMailKey(mail, key)))
	if err != nil {
		log.Println(err)
		return
	}
	println("apperantly i sent")
	// redirect page
}

var loginTmpl = `Hello, here is your login link: <a href="/api/login/verify/%v">LINK</a>`

func VerifyLogin(s *server.Server, w http.ResponseWriter, r *http.Request) {
	mail, outsideKey, err := users.DecodeMailKey(r.URL.Path[len("/api/login/verify/"):])
	if err != nil {
		log.Println(err)
		return
	}

	key, err := s.Users.GetVerify(mail)
	if err != nil {
		log.Println(err)
		return
	}

	if outsideKey == key {
		err = generateSession(s, w, mail)
		if err != nil {
			log.Println(err)
			return
		}
		http.Redirect(w, r, "/", 307)
		return
	}
	http.Error(w, "server error", 502)
}

func generateSession(s *server.Server, w http.ResponseWriter, mail string) error {
	key, err := users.GenerateSessionKey()
	if err != nil {
		return err
	}

	storeToCookie(w, mail, key)
	return s.Users.StoreSession(mail, key)
}

func storeToCookie(w http.ResponseWriter, mail, key string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  users.EncodeMailKey(mail, key),
		Path:   "/",
		MaxAge: 15552000,
	})
}
