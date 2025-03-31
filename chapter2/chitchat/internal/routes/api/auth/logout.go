package api

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api/auth"
)

func SetLogoutApiRoutes(r *mux.Router, cache *database.Cache) {
	logoutRoutes := r.PathPrefix("/logout").Subrouter()

	logoutHandler := auth.NewLogoutHandlerAPI(cache)

	logoutRoutes.HandleFunc("", logoutHandler.LogoutHandlerAPI)
}
