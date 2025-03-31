package ui

import (
	"log"
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/templates"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	parser := templates.NewParseTemplate(
		templates.TemplateConfig{
			Writer: w,
			Name:   "layout",
			Data:   nil,
			Path: []string{
				"views/layout.html",
				"views/public.navbar.html",
				"views/signup.html",
			},
		})

	if err := parser.ParseAndRender(); err != nil {
		log.Println("Error Parsing And Rendering Template: ", err)
		http.Error(w, "Internal Server Error: ", http.StatusInternalServerError)
	}
}
