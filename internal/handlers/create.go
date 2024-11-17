package handlers

import (
	"fmt"
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/service"
)

func (h *ArchiveHandler) CreateArchiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files[]"]
	if len(files) == 0 {
		http.Error(w, "No files provided", http.StatusBadRequest)
		return
	}

	for _, file := range files {
		if !service.IsValidMimeType(file.Header.Get("Content-Type")) {
			http.Error(w, fmt.Sprintf("Invalid file type: %s", file.Filename), http.StatusBadRequest)
			return
		}
	}

	zipData, err := h.archiveService.CreateArchive(files)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating archive: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=archive.zip")
	w.Write(zipData)
}
