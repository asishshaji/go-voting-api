package main

import (
	"fmt"
	"net/http"

	"github.com/asishshaji/go-voting-api/go/src/github.com/asishshaji/go-voting-api/internal/config"
	"github.com/asishshaji/go-voting-api/go/src/github.com/asishshaji/go-voting-api/internal/database"
	"github.com/asishshaji/go-voting-api/go/src/github.com/asishshaji/go-voting-api/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "healthy")
	})

	db, err := database.NewDB(config.NewConfig())

	if err != nil {
		panic(err)
	}

	user.RegisterRoutes(db, router)

	fmt.Println("Starting server...")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error staring server: %s", err)
	}
}
