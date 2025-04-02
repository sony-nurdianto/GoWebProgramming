package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/middleware"
)

func SetThreadRoutesAPI(r *mux.Router, d *database.Database, c *database.Cache) {
	threadRouter := r.PathPrefix("/thread").Subrouter()

	newThreadHandler := api.NewThreadHandlerApi(d, c)

	newAuthMw := middleware.NewMiddleWareAuth(c)
	threadRouter.Handle("/create", middleware.WraperMiddleware(http.HandlerFunc(newThreadHandler.CreateThread), newAuthMw.AuthMiddleware)).Methods(http.MethodPost)
}
