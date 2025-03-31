package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	cr "github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository/cache"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/auth"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/cache"
)

type LogoutHandlerApi struct {
	cache *database.Cache
}

func NewLogoutHandlerAPI(cache *database.Cache) *LogoutHandlerApi {
	return &LogoutHandlerApi{
		cache: cache,
	}
}

func (lh *LogoutHandlerApi) LogoutHandlerAPI(w http.ResponseWriter, r *http.Request) {
	cookies, err := r.Cookie("session_token")
	if err != nil {
		return
	}

	sessionRepo := cr.NewSessionRepo(lh.cache)
	sessionService := cache.NewSessionService(sessionRepo)

	logoutService := auth.NewLogoutService(sessionService)
	if err := logoutService.Logout(cookies.Value); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)

		return
	}

	expiredCookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, &expiredCookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
