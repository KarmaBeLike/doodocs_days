package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
)

func (h *ArchiveHandler) GetArchiveInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, errors.ErrMethodNotAllowed.Message, errors.ErrMethodNotAllowed.Code)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v\n", err)
		http.Error(w, errors.ErrInvalidFile.Message, errors.ErrInvalidFile.Code)
		return
	}
	defer file.Close()

	log.Printf("Received file: %s, size: %d bytes\n", header.Filename, header.Size)

	archiveInfo, err := h.archiveService.GetArchiveInfo(file, header)
	if err != nil {
		if archiveErr, ok := err.(*errors.ErrorResponse); ok {
			log.Printf("Error retrieving archive info: %v\n", archiveErr)
			http.Error(w, archiveErr.Message, archiveErr.Code)
			return
		}
		log.Printf("Unknown error while retrieving archive info: %v\n", err)
		http.Error(w, errors.ErrInternal.Message, errors.ErrInternal.Code)
		return
	}

	log.Printf("Successfully retrieved archive info: %v\n", archiveInfo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(archiveInfo); err != nil {
		log.Printf("Error encoding archive info to JSON: %v\n", err)
		http.Error(w, errors.ErrInternal.Message, errors.ErrInternal.Code)
	}
}
