package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api/auth"
)

func SetLoginAPIRoutes(r *mux.Router, data *database.Database, cache *database.Cache) {
	loginRoutes := r.PathPrefix("/login").Subrouter()

	loginHandler := auth.NewLoginHandlerAPi(data, cache)

	loginRoutes.HandleFunc("/authenticate", loginHandler.AuthenticateLogin).Methods(http.MethodPost)
}
