package ui

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui"
)

func SetLogingRoutes(r *mux.Router) {
	loginRoutes := r.PathPrefix("/login").Subrouter()

	loginRoutes.HandleFunc("", ui.Login)
}
