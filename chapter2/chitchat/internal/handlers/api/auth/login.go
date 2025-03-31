package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/auth"
)

type LoginHandlerApi struct {
	data  *database.Database
	cache *database.Cache
}

func NewLoginHandlerAPi(data *database.Database, cache *database.Cache) *LoginHandlerApi {
	return &LoginHandlerApi{
		data:  data,
		cache: cache,
	}
}

func (lh *LoginHandlerApi) AuthenticateLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Failed To Parse Form: ", err)
		http.Error(w, fmt.Sprintf("failed parse login form: %v", err), http.StatusInternalServerError)
		return
	}

	userRepo := repository.NewUserRepository(lh.data)
	userService := service.NewUserService(userRepo)
	loginService := auth.NewLoginUpService(userService, lh.cache)

	token, err := loginService.AuthenticateLogin(r.PostFormValue("email"), r.PostFormValue("password"))
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrMarshallSession):
			log.Println(err)
			http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)
			return
		case errors.Is(err, auth.ErrSetSession):
			log.Println(err)
			http.Error(w, fmt.Sprintf("authenticated failed: %v", err), http.StatusInternalServerError)
			return
		case errors.Is(err, auth.ErrEncryptToken):
			log.Println(err)
			http.Error(w, fmt.Sprintf("internal server error failed: %v", err), http.StatusInternalServerError)
			return
		case errors.Is(err, auth.ErrUnAuthorizedUser):
			log.Println("unauthorized access deniend wrong password")
			http.Error(w, fmt.Sprintf("%v", err), http.StatusUnauthorized)
			return
		case errors.Is(err, auth.ErrUserNotFound):
			log.Println("user not found")
			http.Error(w, "User Not Found", http.StatusNotFound)
			return
		default:
			log.Println(err)
			http.Error(w, fmt.Sprintf("internal server error: %v", err), http.StatusInternalServerError)
			return
		}
	}

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
