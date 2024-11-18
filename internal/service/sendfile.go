package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/scorredoira/email"
)

// Отправка файла по email
func (s *ArchiveService) SendFile(file multipart.File, fileHeader *multipart.FileHeader, emails []string) error {
	// Check if the file MIME type is valid
	if !isValidFileType(fileHeader) {
		return fmt.Errorf("invalid  mimetype:%v", fileHeader.Filename)
	}

	// Create a temp directory
	tempDir := os.TempDir()

	filePath := filepath.Join(tempDir, fileHeader.Filename)

	// Create a new file in a tempDir
	destination, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Copy the file content into the created file
	_, err = io.Copy(destination, file)
	if err != nil {
		return err
	}

	// Send email with the file attachment
	err = sendFileToEmails(filePath, emails)
	if err != nil {
		return err
	}
	return nil
}

func sendFileToEmails(filePath string, emails []string) error {
	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	if from == "" || pass == "" || host == "" || port == "" {
		return errors.New("SMTP configuration is missing")
	}
	if len(emails) == 0 {
		return errors.New("no recipient email addresses provided")
	}
	// Set up authentication information.
	auth := smtp.PlainAuth("", from, pass, host)

	// Create an email message
	m := email.NewMessage("Subject: File submission", "Please see a file attached bellow")
	m.To = emails

	// Add attachments
	if err := m.Attach(filePath); err != nil {
		return err
	}

	if err := email.Send("smtp.gmail.com:587", auth, m); err != nil {
		return err
	}

	return nil
}

// Check if the file MIME type is valid
func isValidFileType(fileHeader *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/pdf": true,
	}
	return allowedTypes[fileHeader.Header.Get("Content-Type")]
}
