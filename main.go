package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yogesh-p-thakare3110/go-voting-api/internal/config"
	"github.com/yogesh-p-thakare3110/go-voting-api/internal/database"
	"github.com/yogesh-p-thakare3110/go-voting-api/user"
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

	fmt.Println("Server Start...")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error staring server: %s", err)
	}
}
