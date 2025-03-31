package ui

import (
	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui"
)

func SetIndexRoutes(r *mux.Router, data *database.Database) {
	indexRoutes := r.PathPrefix("/").Subrouter()

	newIndexHandler := ui.NewIndexHandlerUI(data)
	indexRoutes.HandleFunc("/", newIndexHandler.Index)
}
