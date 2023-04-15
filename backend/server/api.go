package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	*Service
}

func SetupRoutes(h *Handler, router *mux.Router) {
	/* -------------- Index -------------- */
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello world")
	})

	/* -------------- Trips -------------- */
	router.HandleFunc("/create", h.CreateUser).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
}

func (h *Handler) CreateUser(res http.ResponseWriter, req *http.Request) {
	// Set response to JSON format
	res.Header().Set("Content-Type", "application/json")
	// Get request body and map it to a user
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}
	// Create user
	created_user, err := h.Service.CreateUser(&user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	// Transform created user into JSON
	response, err := json.Marshal(created_user)
	if err != nil {

		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))
}

func (h *Handler) Login(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var login Login
	err := json.NewDecoder(req.Body).Decode(&login)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "Error unmarshalling"}`))
		return
	}
	result, err := h.Service.Login(&login)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	response, err := json.Marshal(result)
	if err != nil {

		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling"}`))
		return
	}

	// Send user as JSON
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(response))
}

type Login struct {
	Email    string
	Password string
}
