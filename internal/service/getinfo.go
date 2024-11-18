package service

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/KarmaBeLike/doodocs_days/internal/entities"
	"github.com/KarmaBeLike/doodocs_days/internal/errors"
)

func (s *ArchiveService) GetArchiveInfo(file io.Reader, header *multipart.FileHeader) (*entities.ArchiveInfo, error) {
	log.Printf("Starting to process archive: %s\n", header.Filename)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file content: %s\n", err)
		return nil, errors.ErrFileRead
	}

	zipReader, err := zip.NewReader(bytes.NewReader(fileBytes), int64(header.Size))
	if err != nil {
		log.Printf("Error creating zip reader for file %s: %s\n", header.Filename, err)
		return nil, errors.ErrNotArchive
	}

	var totalSize float64
	var files []entities.FileDetails

	for _, f := range zipReader.File {
		fileInfo := f.FileInfo()

		mimeType := mime.TypeByExtension(strings.ToLower(filepath.Ext(f.Name)))
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}

		size := float64(fileInfo.Size())
		totalSize += size

		log.Printf("File found in archive: %s, Size: %f, MIME Type: %s\n", f.Name, size, mimeType)

		files = append(files, entities.FileDetails{
			FilePath: f.Name,
			Size:     size,
			MimeType: mimeType,
		})
	}

	archiveInfo := &entities.ArchiveInfo{
		Filename:    header.Filename,
		ArchiveSize: float64(header.Size),
		TotalSize:   totalSize,
		TotalFiles:  float64(len(files)),
		Files:       files,
	}

	log.Printf("Archive processed successfully. Total files: %d, Total size: %f\n", len(files), totalSize)

	return archiveInfo, nil
}
