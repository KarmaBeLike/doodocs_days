package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
)

func (h *ArchiveHandler) GetArchiveInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		apiErr := errors.ErrInvalidFile
		http.Error(w, apiErr.Message, apiErr.Code)
		return
	}
	defer file.Close()

	archiveInfo, err := h.archiveService.GetArchiveInfo(file, header)
	if err != nil {
		if archiveErr, ok := err.(*errors.ErrorResponse); ok {
			http.Error(w, archiveErr.Message, archiveErr.Code)
			return
		}
		apiErr := errors.ErrInternal
		http.Error(w, apiErr.Message, apiErr.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(archiveInfo); err != nil {
		http.Error(w, errors.ErrInternal.Message, errors.ErrInternal.Code)
	}
}
