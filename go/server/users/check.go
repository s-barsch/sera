package users

import (
	"net/http"
	"fmt"
)

type Auth struct {
	User	   *User
	Subscriber bool
}

func (u *Users) CheckAuth(r *http.Request) (*Auth, error) {
	c, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	mail, outsideKey, err := DecodeMailKey(c.Value)
	if err != nil {
		return nil, err
	}

	key, err := u.GetSessionKey(mail)
	if err != nil {
		return nil, err
	}

	if key != outsideKey {
		return nil, fmt.Errorf("key mismatch")
	}

	user, err := u.LookupUser(mail)
	if err != nil {
		return nil, err
	}
	
	return &Auth{
		User:	    user,
		Subscriber: isSubscriber(user),
	}, nil
}

func noAuth() *Auth {
	return nil
}

func isSubscriber(u *User) bool {
	return u.PaypalID != ""
}


