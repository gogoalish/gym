package service

import (
	"project/internal/model"
)

type auth struct {
	userRepo model.UserRepository
	sess     model.SessionService
}

func NewAuthService(userRepo model.UserRepository) model.AuthService {
	return &auth{
		userRepo: userRepo,
		// sessionReop: sessRepo,
	}
}

func (a *auth) LogIn() {
}

func (a *auth) LogOut() {
}

func (a *auth) SignUp(user *model.User) {
}
