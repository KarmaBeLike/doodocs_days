package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
)

var allowedMimes = map[string]bool{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
	"application/xml": true,
	"image/jpeg":      true,
	"image/png":       true,
}

func IsValidMimeType(mimeType string) bool {
	return allowedMimes[mimeType]
}

func (s *ArchiveService) CreateArchive(files []*multipart.FileHeader) ([]byte, error) {
	log.Println("Starting to create archive with files")

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, fileHeader := range files {
		log.Printf("Processing file: %s\n", fileHeader.Filename)
		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("Error opening file %s: %s\n", fileHeader.Filename, err)
			return nil, fmt.Errorf("%w: %s", errors.ErrFileOpenFailed, fileHeader.Filename)
		}
		defer file.Close()

		if !IsValidMimeType(fileHeader.Header.Get("Content-Type")) {
			log.Printf("Invalid MIME type for file %s\n", fileHeader.Filename)
			return nil, fmt.Errorf("%w: %s", errors.ErrInvalidMime, fileHeader.Filename)
		}

		zipFileWriter, err := zipWriter.Create(fileHeader.Filename)
		if err != nil {
			log.Printf("Error creating zip entry for file %s: %s\n", fileHeader.Filename, err)
			return nil, fmt.Errorf("%w: %s", errors.ErrZipCreation, fileHeader.Filename)
		}

		if _, err := io.Copy(zipFileWriter, file); err != nil {
			log.Printf("Error copying file %s to zip: %s\n", fileHeader.Filename, err)
			return nil, fmt.Errorf("%w: %s", errors.ErrZipWriteFailed, fileHeader.Filename)
		}
	}

	if err := zipWriter.Close(); err != nil {
		log.Printf("Error closing zip writer: %s\n", err)
		return nil, errors.ErrZipCloseFailed
	}
	log.Println("Archive created successfully")

	return buf.Bytes(), nil
}
