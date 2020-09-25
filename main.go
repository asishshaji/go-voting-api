package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Healthy")
	})
	fmt.Println("Starting server")

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		fmt.Printf("Error starting server : %s", err)
	}

}
