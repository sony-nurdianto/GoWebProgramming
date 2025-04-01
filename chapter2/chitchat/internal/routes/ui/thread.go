package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/middleware"
)

func SetThreadRoutesUi(r *mux.Router, c *database.Cache) {
	threadRouter := r.PathPrefix("/thread").Subrouter()

	newAuthMw := middleware.NewMiddleWareAuth(c)
	threadRouter.Handle("/new", middleware.WraperMiddleware(http.HandlerFunc(ui.ThreadFormHandlerUi), newAuthMw.AuthMiddleware))
}
