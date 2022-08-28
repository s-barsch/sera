package auth

import (
	"sacer/go/server"
	"sacer/go/server/users"
	"net/http"
	"fmt"
	"encoding/json"
	"io"
	"log"
	"encoding/base64"
	"strings"
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

	http.SetCookie(w, &http.Cookie{
		Name:   "hello",
		Value:  "hai",
		Path:   "/",
		MaxAge: 15552000,
	})

	key, err:= generateLoginKey()
	if err != nil {
		log.Println(err)
		return 
	}

	err = s.Users.StoreVerify(mail, key)
	if err != nil {
		log.Println(err)
		return 
	}

	err = send(user.Mail, fmt.Sprintf(loginTmpl, encodeMailKey(mail, key)))
	if err != nil {
		log.Println(err)
		return 
	}
	println("apperantly i sent")
	// redirect page
}

func generateLoginKey() (string, error) {
	return GenerateRandomStringURLSafe(32)
}

var loginTmpl = `Hello, here is your login link: <a href="/api/login/verify/%v">LINK</a>`

func encodeMailKey(mail, key string) string {
	return base64.StdEncoding.EncodeToString([]byte(mail + "+" + key))
}

func decodeMailKey(b64 string) (mail, key string, err error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", "", err
	}
	str := string(b)
	i := strings.Index(str, "+")
	if i <= 0 {
		err = fmt.Errorf("decodeVerify: cannot get key and string")
		return "", "", err
	}
	return str[:i], str[i+1:], nil

}

func VerifyLogin(s *server.Server, w http.ResponseWriter, r *http.Request) {
	mail, outsideKey, err := decodeMailKey(r.URL.Path[len("/api/login/verify/"):])
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
	key, err := generateSessionKey()
	if err != nil {
		return err
	}

	storeToCookie(w, mail, key)
	return s.Users.StoreSession(mail, key)
}

func generateSessionKey() (string, error) {
	return GenerateRandomStringURLSafe(40)
}

func storeToCookie(w http.ResponseWriter, mail, key string) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  encodeMailKey(mail, key),
		Path:   "/",
		MaxAge: 15552000,
	})
}

