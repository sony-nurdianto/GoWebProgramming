package threads

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/templates"
)

type ThreadsHandlerUi struct {
	db *database.Database
}

func NewThreadsHandlerUi(data *database.Database) *ThreadsHandlerUi {
	return &ThreadsHandlerUi{
		db: data,
	}
}

func (th *ThreadsHandlerUi) ThreadsDetailsHandlerUI(w http.ResponseWriter, r *http.Request) {
	threadUUID := r.URL.Query().Get("threadUUID")
	threadId := r.URL.Query().Get("threadId")

	newThreadRepo := repository.NewThreadRepository(th.db)
	newThreadService := service.NewThreadService(newThreadRepo)

	thread, err := newThreadService.GetThreadDetails(threadUUID)
	if err != nil {
		log.Println("Failed TO Get Data Threads Details")
		http.Error(w, fmt.Sprintf("Internal Server Error : %v", err), http.StatusInternalServerError)
		return
	}

	threadDetailsUiData := models.ContentData{
		Thread:    thread,
		BtnAddUrl: fmt.Sprintf("/post/new?threadUUID=%s&threadId=%s", threadUUID, threadId),
		IsThread:  false,
	}

	parser := templates.NewParseTemplate(
		templates.TemplateConfig{
			Writer: w,
			Name:   "layout",
			Data:   threadDetailsUiData,
			Path: []string{
				"views/private/layout.html",
				"views/private/navbar.html",
				"views/private/thread_details.html",
				"views/private/component/addButton.html",
			},
		})

	if err := parser.ParseAndRender(); err != nil {
		log.Println("Error Parsing And Rendering Template: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
		return
	}
}
