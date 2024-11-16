package service

import (
	"archive/zip"
	"bytes"
	"io"
	"mime"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/KarmaBeLike/doodocs_days/internal/entities"
	"github.com/KarmaBeLike/doodocs_days/internal/errors"
)

func GetArchiveInfo(file io.Reader, header *multipart.FileHeader) (*entities.ArchiveInfo, error) {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.ErrFileRead
	}

	zipReader, err := zip.NewReader(bytes.NewReader(fileBytes), int64(header.Size))
	if err != nil {
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

		files = append(files, entities.FileDetails{
			FilePath: f.Name,
			Size:     size,
			MimeType: mimeType,
		})
	}
	return &entities.ArchiveInfo{
		Filename:    header.Filename,
		ArchiveSize: float64(header.Size),
		TotalSize:   totalSize,
		TotalFiles:  float64(len(files)),
		Files:       files,
	}, nil
}
