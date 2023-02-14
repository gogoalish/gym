package routes

import (
	"database/sql"
	"log"
	"net/http"

	"project/internal/handlers"
	"project/internal/repository"
	"project/internal/service"
)

const (
	LogIn  = "/login"
	LogOut = "/logout"
	SignUp = "/signup"
)

func InitRoutes(db *sql.DB, mux *http.ServeMux) {
	// repository
	log.Println("init user repositories")
	userRepo := repository.NewUserRepository(db)
	// sessionRepo := repository.NewSessionRepository(db)
	// postRepo := repository.NewPostRepository(db)
	// categoryRepo := repository.NewCategoryRepository(db)
	// commentRepo := repository.NewCommentRepository(db)
	// reactionRepo := repository.NewReactionRepository(db)

	// service
	log.Println("init authentication service")
	authServ := service.NewAuthService(userRepo)

	// handlers
	log.Println("init authentication handlers")
	authHandlers := handlers.NewAuthHandlers(authServ)

	// middleware
	log.Println("init middleware handlers")
	middleware := handlers.NewMiddlewareHandlers()

	// handlefunc
	mux.HandleFunc(LogIn, authHandlers.LogIn)
	mux.HandleFunc(LogOut, authHandlers.LogOut)
	mux.HandleFunc(SignUp, authHandlers.SignUp)

	mux.Handle("/", middleware.PanicRecover(mux))
	// handle(get, address/qwer, ShowPRofile)
	// handle(post, address/qwer, CreateUser)
}
