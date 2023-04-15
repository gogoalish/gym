package server

import (
	"errors"
	"fmt"
)

type Service struct {
	UserRepo *Repo
}

func (s *Service) CreateUser(user *User) (*User, error) {
	if user == nil {
		err := errors.New("User is empty")
		return nil, err
	}

	if user.Name == "" || user.Email == "" {
		err := errors.New("missing required fields")
		return nil, err
	}
	// FMT
	fmt.Println(user)
	err := s.UserRepo.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (s *Service) Login(login *Login) (*User, error) {
	if login == nil {
		err := errors.New("form is empty")
		return nil, err
	}

	if login.Email == "" || login.Password == "" {
		err := errors.New("missing required fields")
		return nil, err
	}
	// FMT
	fmt.Println(login)
	user, err := s.UserRepo.Login(login)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("No users found")
	}
	return &user, nil
}
