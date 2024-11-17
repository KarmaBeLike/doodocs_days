package handlers

import "github.com/KarmaBeLike/doodocs_days/internal/service"

type ArchiveHandler struct {
	archiveService *service.ArchiveService
}

func NewArchiveHandler(service *service.ArchiveService) *ArchiveHandler {
	return &ArchiveHandler{archiveService: service}
}
