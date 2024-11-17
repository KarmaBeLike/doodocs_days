package routers

import (
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/handlers"
	"github.com/KarmaBeLike/doodocs_days/internal/service"
)

func SetupRouters() *http.ServeMux {
	mux := http.NewServeMux()

	archiveService := &service.ArchiveService{}
	archiveHandler := handlers.NewArchiveHandler(archiveService)

	mux.HandleFunc("/archive/info", archiveHandler.GetArchiveInfoHandler)
	mux.HandleFunc("/archive/files", archiveHandler.CreateArchiveHandler)

	return mux
}
