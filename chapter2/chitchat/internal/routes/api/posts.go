package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/middleware"
)

func SetPostRoutesAPI(r *mux.Router, d *database.Database, c *database.Cache) {
	postRouter := r.PathPrefix("/post").Subrouter()

	newPostHandler := api.NewPostHandlerApi(d)

	newAuthMw := middleware.NewMiddleWareAuth(c)
	postRouter.Handle("/new", middleware.WraperMiddleware(http.HandlerFunc(newPostHandler.PostComment), newAuthMw.AuthMiddleware)).Methods(http.MethodPost)
}
