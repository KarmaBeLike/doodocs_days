package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
	"github.com/KarmaBeLike/doodocs_days/internal/service"
)

func (h *ArchiveHandler) CreateArchiveHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		log.Printf("Method not allowed: %s\n", r.Method)
		http.Error(w, errors.ErrMethodNotAllowed.Message, errors.ErrMethodNotAllowed.Code)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing multipart form: %v\n", err)
		http.Error(w, errors.ErrFormParsingFailed.Message, errors.ErrFormParsingFailed.Code)
		return
	}

	files := r.MultipartForm.File["files[]"]
	if len(files) == 0 {
		log.Println("No files provided in the form")
		http.Error(w, errors.ErrNoFileProvided.Message, errors.ErrNoFileProvided.Code)
		return
	}

	log.Printf("Received %d file(s)\n", len(files))

	for _, file := range files {
		if !service.IsValidMimeType(file.Header.Get("Content-Type")) {
			log.Printf("Invalid file type: %s\n", file.Filename)
			http.Error(w, fmt.Sprintf("Invalid file type: %s", file.Filename), http.StatusBadRequest)
			return
		}
		log.Printf("File validated: %s\n", file.Filename)
	}

	zipData, err := h.archiveService.CreateArchive(files)
	if err != nil {
		log.Printf("Error creating archive: %v\n", err)
		http.Error(w, fmt.Sprintf("Error creating archive: %v", err), http.StatusInternalServerError)
		return
	}
	log.Println("Successfully created the archive, sending response")

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=archive.zip")
	w.Write(zipData)
}
