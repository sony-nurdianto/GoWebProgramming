package middleware

import (
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers"
)

func DynamicPrefix(next http.Handler) http.Handler {
	stripPrefix := handlers.NewPrefixHandler(next)

	return http.HandlerFunc(stripPrefix.StripDynamicPrefix)
}
