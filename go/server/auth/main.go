package auth

import (
	"net/http"
)

type Auth struct {
	Subscriber bool
}

func CheckAuth(r *http.Request) *Auth {
	return &Auth{
		Subscriber: false,
	}
}
