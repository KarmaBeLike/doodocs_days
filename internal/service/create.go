package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
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
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		if !IsValidMimeType(fileHeader.Header.Get("Content-Type")) {
			return nil, fmt.Errorf("invalid file type: %s", fileHeader.Filename)
		}

		zipFileWriter, err := zipWriter.Create(fileHeader.Filename)
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(zipFileWriter, file); err != nil {
			return nil, err
		}
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
