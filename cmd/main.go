package main

import (
	"log"

	"github.com/KarmaBeLike/doodocs_days/internal/routers"
)

func main() {
	router := routers.SetupRouters()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
