package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	api "github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/routes/api/auth"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/routes/ui"
)

func main() {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := database.NewDatabase(dburl)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cache, err := database.NewCache()
	if err != nil {
		log.Fatal(err)
	}

	defer cache.Close()

	router := mux.NewRouter()

	path, err := filepath.Abs("public")
	log.Println(path)
	if err != nil {
		log.Fatal(err)
	}

	files := http.FileServer(http.Dir(path))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", files))

	ui.SetIndexRoutes(router, db, cache)
	ui.SetLogingRoutes(router)
	ui.SetSignUpRoutes(router)
	api.SetSignUpAPIRoutes(router, db)
	api.SetLoginAPIRoutes(router, db, cache)
	api.SetLogoutApiRoutes(router, cache)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server is running at %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-stop
	log.Println("shuting Down Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Forced Shutdown %v", err)
	}

	log.Println("Server exited gracefully")
}
