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

type ThreadDetailsUiData struct {
	Threads   models.Thread
	BtnAddUrl string
}

type ThreadsHandlerUi struct {
	db *database.Database
}

func NewThreadsHandlerUi(data *database.Database) *ThreadsHandlerUi {
	return &ThreadsHandlerUi{
		db: data,
	}
}

func (th *ThreadsHandlerUi) ThreadsDetailsHandlerUI(w http.ResponseWriter, r *http.Request) {
	threadUUID := r.URL.Query().Get("id")

	newThreadRepo := repository.NewThreadRepository(th.db)
	newThreadService := service.NewThreadService(newThreadRepo)

	thread, err := newThreadService.GetThreadDetails(threadUUID)
	if err != nil {
		log.Println("Failed TO Get Data Threads Details")
		http.Error(w, fmt.Sprintf("Internal Server Error : %v", err), http.StatusInternalServerError)
		return
	}

	threadDetailsUiData := ThreadDetailsUiData{
		Threads:   thread,
		BtnAddUrl: "/thread/comment?id=" + threadUUID,
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
