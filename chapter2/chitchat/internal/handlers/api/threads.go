package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/encryption"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
	cr "github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository/cache"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/cache"
)

type ThreadHandlerApi struct {
	data  *database.Database
	cache *database.Cache
}

func NewThreadHandlerApi(data *database.Database, cache *database.Cache) *ThreadHandlerApi {
	return &ThreadHandlerApi{
		data:  data,
		cache: cache,
	}
}

func (th *ThreadHandlerApi) CreateThread(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Failed to ParseForm: ", err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %s", err), http.StatusInternalServerError)
		return
	}

	topic := r.FormValue("topic")

	if topic == "" {
		log.Println("Topic is empty")
		http.Error(w, "Bad Request: Topic is empty", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println("Failed to Get Cookie: ", err)
		http.Error(w, fmt.Sprintf("Unauthorized Token is missing: %s", err), http.StatusUnauthorized)
		return
	}

	token, err := encryption.VerifyWebToken(cookie.Value)
	if err != nil {
		if errors.Is(err, encryption.ErrDecryptFailed) {
			log.Printf("Access Denied: %v\n", err)
			http.Error(w, "Access Denied Failed To Decrypt Token", http.StatusUnauthorized)
			return
		}

		log.Println("Failed Decrypt Token")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sessionRepo := cr.NewSessionRepo(th.cache)
	sessionService := cache.NewSessionService(sessionRepo)

	sessionData, err := sessionService.GetSession(token.Subject)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Internal Servr Error: %v", err), http.StatusInternalServerError)
		return
	}

	newThreadRepo := repository.NewThreadRepository(th.data)
	newThreadService := service.NewThreadService(newThreadRepo)

	_, err = newThreadService.CreateThread(topic, sessionData.Id)
	if err != nil {
		log.Println("Failed To Create Thread: ", err)
		http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
