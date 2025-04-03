package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/ui/threads"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/middleware"
)

func SetThreadRoutesUi(r *mux.Router, d *database.Database, c *database.Cache) {
	threadRouter := r.PathPrefix("/thread").Subrouter()

	newAuthMw := middleware.NewMiddleWareAuth(c)
	threadRouter.Handle("/new",
		middleware.WraperMiddleware(http.HandlerFunc(threads.ThreadFormHandlerUi),
			newAuthMw.AuthMiddleware)).Methods(http.MethodGet)

	newThreadHandler := threads.NewThreadsHandlerUi(d)

	threadRouter.Handle("/read",
		middleware.WraperMiddleware(
			http.HandlerFunc(newThreadHandler.ThreadsDetailsHandlerUI),
			newAuthMw.AuthMiddleware)).Methods(http.MethodGet)
}
