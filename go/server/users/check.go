package users

import (
	"net/http"
	"fmt"
)

type Auth struct {
	User	   *User
	Subscriber bool
}

func (a *Auth) SignedIn() bool {
	return a.User.Mail != ""
}

func (u *Users) CheckAuth(r *http.Request) (*Auth, error) {
	na := noAuth()
	c, err := r.Cookie("session")
	if err != nil {
		return na, err
	}

	mail, outsideKey, err := DecodeMailKey(c.Value)
	if err != nil {
		return na, err
	}

	key, err := u.GetSessionKey(mail)
	if err != nil {
		return na, err
	}

	if key != outsideKey {
		return na, fmt.Errorf("key mismatch")
	}

	user, err := u.LookupUser(mail)
	if err != nil {
		return na, err
	}

	sub := isSubscriber(user)
	user.PaypalID = ""
	
	return &Auth{
		User:       user,
		Subscriber: sub,
	}, nil
}

func noAuth() *Auth {
	return &Auth{
		User: &User{},
	}
}

func isSubscriber(u *User) bool {
	return u.PaypalID != ""
}


