package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/middleware"
)

func SetIndexRoutes(r *mux.Router, data *database.Database, cache *database.Cache) {
	indexRoutes := r.PathPrefix("/").Subrouter()

	newIndexHandler := ui.NewIndexHandlerUI(data)

	indexRoutes.HandleFunc("/", newIndexHandler.Index).Methods(http.MethodGet)

	// protected routes
	newAuthMw := middleware.NewMiddleWareAuth(cache)
	indexRoutes.Handle("/home", middleware.WraperMiddleware(http.HandlerFunc(newIndexHandler.Home), newAuthMw.AuthMiddleware)).Methods(http.MethodGet)
}
