package service

import (
	"io"
	"log"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
	"github.com/scorredoira/email"
)

// Отправка файла по email
func (s *ArchiveService) SendFile(file multipart.File, fileHeader *multipart.FileHeader, emails []string) error {
	log.Println("Starting to send file:", fileHeader.Filename)

	if !isValidFileType(fileHeader) {
		log.Printf("Invalid file type: %s\n", fileHeader.Filename)
		return errors.ErrInvalidMime
	}

	// Временная директория для сохранения файла
	tempDir := os.TempDir()

	filePath := filepath.Join(tempDir, fileHeader.Filename)

	log.Printf("Saving file to temporary directory: %s\n", filePath)

	// Create a new file in a tempDir
	destination, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error saving file: %s\n", err)
		return errors.ErrFileSaveFailed
	}
	defer destination.Close()

	// Copy the file content into the created file
	_, err = io.Copy(destination, file)
	if err != nil {
		log.Printf("Error copying file content: %s\n", err)
		return errors.ErrFileSaveFailed
	}

	log.Println("Sending file to emails:", emails)

	err = sendFileToEmails(filePath, emails)
	if err != nil {
		log.Printf("Error sending file to emails: %s\n", err)
		return err
	}
	log.Printf("File %s sent successfully\n", fileHeader.Filename)
	return nil
}

func sendFileToEmails(filePath string, emails []string) error {
	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	if from == "" || pass == "" || host == "" || port == "" {
		log.Println("SMTP configuration is missing")
		return errors.ErrSMTPConfigMissing
	}
	if len(emails) == 0 {
		log.Println("No recipient emails provided")
		return errors.ErrNoEmailsProvided
	}
	// Set up authentication information.
	auth := smtp.PlainAuth("", from, pass, host)

	// Create an email message
	m := email.NewMessage("Subject: File submission", "Please see a file attached bellow")
	m.To = emails

	// Add attachments
	if err := m.Attach(filePath); err != nil {
		log.Printf("Error attaching file: %s\n", err)
		return errors.ErrFileSendFailed
	}
	log.Printf("Sending email to %v\n", emails)
	if err := email.Send(host+":"+port, auth, m); err != nil {
		log.Printf("Error sending email: %s\n", err)
		return errors.ErrEmailSendFailed
	}
	log.Printf("Email sent to %v with attachment %s\n", emails, filePath)
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
