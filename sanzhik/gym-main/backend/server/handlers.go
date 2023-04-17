package server

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Handler struct {
	tmpl    *template.Template
	Service *Service
	logger  *Logger
	// cfg       config
	validator *Validator
	// wg        sync.WaitGroup
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		// tmpl:      template.Must(template.ParseGlob("./templates/*")),
		Service:   service,
		validator: NewValidator(),
		logger:    New(os.Stdout, LevelInfo),
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.ExecuteTemplate(w, "show_posts.html", nil)
	if err != nil {
		h.logger.PrintError(err, nil)
		h.serverErrorResponse(w, r, err)
	}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/signup" {
		h.logger.PrintError(fmt.Errorf("handler: signup: not found"), nil)
		h.notFoundResponse(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err := h.tmpl.ExecuteTemplate(w, "signup.html", nil)
		if err != nil {
			h.logger.PrintError(err, nil)
			h.serverErrorResponse(w, r, err)
			return
		}
	case http.MethodPost:
		var input struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := readJSON(w, r, &input)
		if err != nil {
			h.badRequestResponse(w, r, err)
			return
		}

		user := &User{
			Name:     input.Name,
			Email:    input.Email,
			Password: Password{Plaintext: input.Password},
		}

		u, err := h.Service.AuthService.Signup(h.validator, user)

		if err == ErrUserExists {
			h.logger.PrintError(err, nil)
			h.editConflictResponse(w, r)
			return
		}

		if err == ErrInternalServer {
			h.logger.PrintError(err, nil)
			h.serverErrorResponse(w, r, err)
			return
		}

		if err == ErrFormValidation {
			h.logger.PrintError(err, nil)
			h.failedValidationResponse(w, r, h.validator.Errors)
			return
		}

		err = writeJSON(w, http.StatusAccepted, envelope{"user": u}, nil)
		if err != nil {
			h.logger.PrintError(err, nil)
			h.serverErrorResponse(w, r, err)
		}

	default:
		h.logger.PrintError(fmt.Errorf("signup: method not allowed"), nil)
		h.methodNotAllowedResponse(w, r)
		return
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/login" {
		fmt.Println("handler.login: not found")
		h.logger.PrintError(fmt.Errorf("handler.login: not found"), nil)
		h.notFoundResponse(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err := h.tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			h.logger.PrintError(err, nil)
			h.serverErrorResponse(w, r, err)
			return
		}
	case http.MethodPost:
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := readJSON(w, r, &input)
		if err != nil {
			h.badRequestResponse(w, r, err)
			return
		}

		user := &User{
			Email:    input.Email,
			Password: Password{Plaintext: input.Password},
		}

		err = h.Service.AuthService.Login(user)
		if err != nil {
			switch err {
			case ErrUserNotFound:
				h.logger.PrintError(fmt.Errorf("handler:login: user not found"), nil)
				h.badRequestResponse(w, r, err)
				return
			case ErrInvalidPassword:
				h.logger.PrintError(fmt.Errorf("handler:login: password is not correct"), nil)
				h.badRequestResponse(w, r, err)
				return
			default:
				h.logger.PrintError(fmt.Errorf("handler:login: password is not correct"), nil)
				h.serverErrorResponse(w, r, err)
				return
			}
		}

		cookie := http.Cookie{}
		cookie.Name = "access_token"
		cookie.Value = *user.Token
		cookie.Expires = *user.Expires
		cookie.Path = "/"
		cookie.HttpOnly = true
		http.SetCookie(w, &cookie)
		writeJSON(w, http.StatusOK, envelope{"token": user.Token}, nil)
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/logout" {
		fmt.Println("handler:logout: not found")
		h.logger.PrintError(fmt.Errorf("handler:logout: not found"), nil)
		h.notFoundResponse(w, r)
		return
	}

	c, err := r.Cookie("access_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			h.logger.PrintError(fmt.Errorf("handler:logout: unauthorized"), nil)
			// writeJSON(w, http.StatusUnauthorized)
			return
		}
		fmt.Println("handler:logout: " + err.Error())
		h.serverErrorResponse(w, r, err)
		return

	}
	err = h.Service.AuthService.DeleteToken(c.Value)
	if err != nil {
		h.logger.PrintError(fmt.Errorf("handler:logout: couldn't delete token"), nil)
		h.serverErrorResponse(w, r, err)
		return
	}

	cookie := &http.Cookie{
		Name:   "access_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
