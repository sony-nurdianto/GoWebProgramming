package ui

import (
	"errors"
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/templates"
)

type IndexHandlerUI struct {
	db *database.Database
}

func NewIndexHandlerUI(data *database.Database) *IndexHandlerUI {
	return &IndexHandlerUI{db: data}
}

func (d *IndexHandlerUI) Index(w http.ResponseWriter, r *http.Request) {
	thredRepo := repository.NewThreadRepository(d.db)
	threadService := service.NewThreadService(thredRepo)

	threads, err := threadService.Threads()
	if err != nil {
		log.Println("Error Get Threds: ", err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}

	navbarPath := "views/private.navbar.html"

	_, err = r.Cookie("session_token")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			navbarPath = "views/public.navbar.html"
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	parser := templates.NewParseTemplate(
		templates.TemplateConfig{
			Writer: w,
			Name:   "layout",
			Data:   threads,
			Path: []string{
				"views/layout.html",
				navbarPath,
				"views/index.html",
			},
		})

	if err := parser.ParseAndRender(); err != nil {
		log.Println("Error Parsing And Rendering Template: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}
}
