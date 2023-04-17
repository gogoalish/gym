package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail        = errors.New("invalid email format")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidUsername     = errors.New("invalid username")
	ErrInvalidUsernameLen  = errors.New("username length out of range 32")
	ErrInvalidUsernameChar = errors.New("invalid username characters")
	ErrInternalServer      = errors.New("internal server error")
	ErrConfirmPassword     = errors.New("password doesn't match")
	ErrUserNotFound        = errors.New("user not found")
	ErrUserExists          = errors.New("user already exists")
	ErrFormValidation      = errors.New("form validation failed")
)

type Service struct {
	AuthService
}

func NewService(repository *Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repository.UserRepository),
	}
}

type AuthService interface {
	Signup(*Validator, *User) (*User, error)
	Login(*User) error
	Logout(*User) error
	ParseToken(token string) (*User, error)
	DeleteToken(token string) error
}

type authService struct {
	ur UserRepository
}

func NewAuthService(userRepo UserRepository) AuthService {
	return &authService{
		ur: userRepo,
	}
}

func (as *authService) Signup(v *Validator, user *User) (*User, error) {
	_, err := as.ur.GetUserByEmail(user.Email)
	if err == nil {
		return nil, ErrUserExists
	}

	if ValidateUser(v, user); !v.Valid() {
		return nil, ErrFormValidation
	}

	err = user.Password.Set(user.Password.Plaintext)
	if err != nil {
		return nil, ErrInternalServer
	}
	// TODO: UNIQUE constraint failed: users.name
	u, err := as.ur.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return nil, ErrInternalServer
	}

	return u, nil
}

func (as *authService) Login(user *User) error {
	u, err := as.ur.GetUserByEmail(user.Email)
	if err != nil {
		return ErrUserNotFound
	}

	ok, err := u.Password.Matches(user.Password.Plaintext)
	if err != nil || !ok {
		return err
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(30 * time.Minute)
	user.Token = &sessionToken
	user.Expires = &expiresAt
	user.ID = u.ID
	err = as.ur.SaveToken(user)
	if err != nil {
		return err
	}

	return nil
}

func (as *authService) Logout(user *User) error {
	u, err := as.ur.GetUserByEmail(user.Email)
	if err != nil {
		return ErrUserNotFound
	}

	user.ID = u.ID

	return as.ur.DeleteToken(*user.Token)
}

func (as *authService) ParseToken(token string) (*User, error) {
	return as.ur.GetUserByToken(token)
}

func (as *authService) DeleteToken(token string) error {
	return as.ur.DeleteToken(token)
}

func ValidateEmail(v *Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(Matches(email, EmailRX), "email", "must be provided a valid email address")
}

func ValidatePasswordPlaintext(v *Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 500 bytes long")
}

func ValidateUser(v *Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	ValidatePasswordPlaintext(v, user.Password.Plaintext)
}
