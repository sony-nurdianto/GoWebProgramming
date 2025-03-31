package ui

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui"
)

func SetIndexRoutes(r *mux.Router, data *database.Database, cache *database.Cache) {
	indexRoutes := r.PathPrefix("/").Subrouter()

	newIndexHandler := ui.NewIndexHandlerUI(data)

	indexRoutes.HandleFunc("/", newIndexHandler.Index)

	// need Impl for users dashboard
	// newAuthMw := middleware.NewMiddleWareAuth(cache)
	// indexRoutes.Handle("/home", middleware.WraperMiddleware(http.HandlerFunc(newIndexHandler.Index), newAuthMw.AuthMiddleware))
}
