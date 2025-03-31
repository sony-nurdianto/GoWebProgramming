package auth

import (
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
)

type SignUpHandlerAPI struct {
	data *database.Database
}

func NewSignUpHandlerAPI(data *database.Database) *SignUpHandlerAPI {
	return &SignUpHandlerAPI{
		data: data,
	}
}

func (d *SignUpHandlerAPI) SignUpAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error Parsring Form", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	userRepo := repository.NewUserRepository(d.data)
	userService := service.NewUserService(userRepo)

	if err := userService.CreateUser(newUser); err != nil {
		log.Println("Failed Create User: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
