package auth

import (
	"net/http"
)

type Auth struct {
	Subscriber bool
}

func CheckAuth(r *http.Request) *Auth {
	return &Auth{
		Subscriber: isSubscriber(r),
	}
}

func isSubscriber(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	return c.Value == "supersecret"
}
