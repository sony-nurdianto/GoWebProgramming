package api

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api/auth"
)

func SetSignUpAPIRoutes(r *mux.Router, data *database.Database) {
	signupRoutes := r.PathPrefix("/signup").Subrouter()

	signupAH := auth.NewSignUpHandlerAPI(data)

	signupRoutes.HandleFunc("/account", signupAH.SignUpAccount)
}
