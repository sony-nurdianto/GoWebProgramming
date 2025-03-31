package ui

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui"
)

func SetSignUpRoutes(r *mux.Router) {
	signupRoutes := r.PathPrefix("/signup").Subrouter()

	signupRoutes.HandleFunc("", ui.Signup)
}
