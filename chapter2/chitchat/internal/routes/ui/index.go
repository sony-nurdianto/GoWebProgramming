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

	newAuthMw := middleware.NewMiddleWareAuth(cache)

	indexRoutes.HandleFunc("/", newIndexHandler.Index)

	indexRoutes.Handle("/home", middleware.WraperMiddleware(http.HandlerFunc(newIndexHandler.Index), newAuthMw.AuthMiddleware))
}
