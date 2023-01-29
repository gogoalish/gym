package routes

import (
	"database/sql"
	"net/http"

	"project/config"
	"project/internal/handlers"
	_authhandlers "project/internal/handlers/auth"

	"project/internal/repository"
	"project/internal/service"
)

const (
	LogIn  = "/login"
	LogOut = "/logout"
	SignUp = "/signup"
)

func InitRoutes(db *sql.DB, mux *http.ServeMux) error {
	err := handlers.ReadTemplate(config.C.P.Template)
	if err != nil {
		return err
	}
	userRepo := repository.NewUserRepository(db)
	authServ := service.NewAuthService(userRepo)
	authHandlers := _authhandlers.NewAuthHandlers(authServ)

	middleware := handlers.NewMiddlewareHandlers()

	mux.HandleFunc(LogIn, authHandlers.LogIn)
	mux.HandleFunc(LogOut, authHandlers.LogOut)
	mux.HandleFunc(SignUp, authHandlers.SignUp)

	

	mux.Handle("", middleware.PanicRecover(mux))
	return nil
}
