package main

import (
	"log"
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/routers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load env: %s", err.Error())
	}

	mux := routers.SetupRouters()

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
