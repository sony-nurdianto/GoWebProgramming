package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api/auth"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
)

type PostHandlerApi struct {
	db *database.Database
}

func NewPostHandlerApi(data *database.Database) *PostHandlerApi {
	return &PostHandlerApi{
		db: data,
	}
}

func (db *PostHandlerApi) PostComment(w http.ResponseWriter, r *http.Request) {
	threadUUID := r.URL.Query().Get("threadUUID")
	threadId := r.URL.Query().Get("threadId")

	intThreadId, err := strconv.Atoi(threadId)
	if err != nil {
		log.Println("Failed to Convert threadId into String: ", err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("Failed to ParseForm: ", err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
		return
	}

	body := r.FormValue("body")

	session := r.Context().Value(auth.ContextKeyUser).(models.Session)

	newPostRepo := repository.NewPostRepository(db.db)
	newPostService := service.NewPostService(newPostRepo)

	_, err = newPostService.PostComment(body, intThreadId, session.UserId)
	if err != nil {
		log.Println("Failed to Create Comment: ", err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/thread/read?threadUUID=%s&threadId=%s", threadUUID, threadId), http.StatusSeeOther)
}
