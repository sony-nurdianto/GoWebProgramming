package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/encryption"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	cr "github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository/cache"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/cache"
)

type UserCtxKey string

const ContextKeyUser UserCtxKey = "user"

type GuardHandler struct {
	handler http.Handler
	cache   *database.Cache
}

func NewGuardHandler(cache *database.Cache, handler http.Handler) *GuardHandler {
	return &GuardHandler{
		handler: handler,
		cache:   cache,
	}
}

func (gh *GuardHandler) AuthGuard(w http.ResponseWriter, r *http.Request) {
	tokenValue := r.Context().Value(ContextKeyUser)
	token, ok := tokenValue.(string)
	if !ok {
		log.Println("Token Not Set or User not signin")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pt, err := encryption.VerifyWebToken(token)
	if err != nil {
		switch {
		case errors.Is(err, encryption.ErrSecretNotSet):
			log.Println("Secret Is Not Set")
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
			return
		case errors.Is(err, encryption.ErrDecryptFailed):
			log.Println("Error verify web token")
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusUnauthorized)
			return
		default:
			log.Printf("Internal Server Error: %v", err)
			http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusUnauthorized)
			return
		}
	}

	if pt.Expiration.Unix() < time.Now().Unix() {
		log.Println("Token expired")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessionRepo := cr.NewSessionRepo(gh.cache)
	sessionService := cache.NewSessionService(sessionRepo)

	sessionData, err := sessionService.GetSession(pt.Subject)
	if err != nil {
		if errors.Is(err, cr.ErrSessionNotFound) {
			log.Println("Error Session Data Not Found")
			http.Error(w, fmt.Sprintf("%v", err), http.StatusUnauthorized)
			return
		}
		log.Println("Error Get Data Session")
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	var session models.Session
	err = json.Unmarshal([]byte(sessionData), &session)
	if err != nil {
		log.Println("Error Get Data Session")
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(r.Context(), ContextKeyUser, session)
	gh.handler.ServeHTTP(w, r.WithContext(ctx))
}
