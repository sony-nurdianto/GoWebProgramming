package threads

import (
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/templates"
)

func ThreadFormHandlerUi(w http.ResponseWriter, r *http.Request) {
	parser := templates.NewParseTemplate(
		templates.TemplateConfig{
			Writer: w,
			Name:   "layout",
			Data:   nil,
			Path: []string{
				"views/private/layout.html",
				"views/private/navbar.html",
				"views/private/component/threadForm.html",
				"views/private/component/addButton.html",
			},
		},
	)

	if err := parser.ParseAndRender(); err != nil {
		log.Println("Error Parsing And Rendering Template: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
