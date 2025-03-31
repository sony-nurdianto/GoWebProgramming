package middleware

import (
	"net/http"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/handlers/api/auth"
)

type MiddleWareAuth struct {
	cache *database.Cache
}

func NewMiddleWareAuth(cache *database.Cache) *MiddleWareAuth {
	return &MiddleWareAuth{
		cache: cache,
	}
}

func (ma *MiddleWareAuth) AuthMiddleware(next http.Handler) http.Handler {
	guardHandler := auth.NewGuardHandler(ma.cache, next)

	return http.HandlerFunc(guardHandler.AuthGuard)
}
