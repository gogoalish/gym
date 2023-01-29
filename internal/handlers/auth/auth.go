package handlers

import (
	"net/http"

	"project/internal/model"
)

type auth struct {
	authServ model.AuthService
}

func NewAuthHandlers(authServ model.AuthService) model.AuthHandlers {
	return &auth{
		authServ: authServ,
	}
}

func (a *auth) LogIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.logInGet(w, r)
	case http.MethodPost:
		a.logInPost(w, r)
	default:
	}
}

func (a *auth) logInGet(w http.ResponseWriter, r *http.Request) {
}

func (a *auth) logInPost(w http.ResponseWriter, r *http.Request) {
}

func (a *auth) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// error
	}
	
}
