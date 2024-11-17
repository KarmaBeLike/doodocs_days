package main

import (
	"log"
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/routers"
)

func main() {
	mux := routers.SetupRouters()

	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
