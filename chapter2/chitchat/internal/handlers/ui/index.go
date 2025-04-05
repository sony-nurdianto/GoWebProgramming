package ui

import (
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
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

func (d *IndexHandlerUI) Home(w http.ResponseWriter, r *http.Request) {
	threadRepo := repository.NewThreadRepository(d.db)
	threadService := service.NewThreadService(threadRepo)

	threads, err := threadService.Threads()
	if err != nil {
		log.Println("Error Get Threds: ", err)
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}

	log.Printf("xxxxxxxxxxxxxxxxxxxxxxx Threads Len %d xxxxxxxxxxxxxxxxxxxxxxx\n", len(threads))

	indexUiData := models.ContentData{
		Threads:   threads,
		BtnAddUrl: "thread/new",
		IsThread:  true,
	}

	parser := templates.NewParseTemplate(
		templates.TemplateConfig{
			Writer: w,
			Name:   "layout",
			Data:   indexUiData,
			Path: []string{
				"views/private/layout.html",
				"views/private/navbar.html",
				"views/private/index.html",
				"views/private/component/addButton.html",
			},
		})

	if err := parser.ParseAndRender(); err != nil {
		log.Println("Error Parsing And Rendering Template: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}
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

	parser := templates.NewParseTemplate(
		templates.TemplateConfig{
			Writer: w,
			Name:   "layout",
			Data:   threads,
			Path: []string{
				"views/public/layout.html",
				"views/public/navbar.html",
				"views/public/index.html",
			},
		})

	if err := parser.ParseAndRender(); err != nil {
		log.Println("Error Parsing And Rendering Template: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}
}
