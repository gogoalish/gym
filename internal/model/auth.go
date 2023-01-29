package model

import "net/http"

type AuthService interface{}

type AuthHandlers interface {
	LogIn(http.ResponseWriter, *http.Request)
	LogOut(http.ResponseWriter, *http.Request)
	SignUp(http.ResponseWriter, *http.Request)
}
